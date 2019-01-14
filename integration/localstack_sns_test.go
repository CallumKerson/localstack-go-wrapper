package integration

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/callumkerredwards/localstack-go-wrapper/localstack"
	"github.com/stretchr/testify/assert"
)

func TestLocalstackSNSStarted(t *testing.T) {
	//given
	//AWS session
	sess, err := session.NewSession(&aws.Config{
		Credentials: testCredentials,
		Region:      testRegion,
		Endpoint:    aws.String("http://localhost:" + localstack.SNSPort),
		DisableSSL:  disableSSL,
	})
	if err != nil {
		t.Errorf("Could not create new AWS session: %v", err)
	}

	//when
	//Create SNS client and call the list topics API
	snsClient := sns.New(sess)
	_, err = snsClient.ListTopics(&sns.ListTopicsInput{})

	//then
	//API did not produce an error
	assert.Nil(t, err, "ListTopics API should not produce an error")
}
