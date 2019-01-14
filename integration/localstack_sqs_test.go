package integration

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/callumkerredwards/localstack-go-wrapper/localstack"
	"github.com/stretchr/testify/assert"
)

func TestLocalstackSQSStarted(t *testing.T) {
	//given
	//AWS session
	sess, err := session.NewSession(&aws.Config{
		Credentials: testCredentials,
		Region:      testRegion,
		Endpoint:    aws.String("http://localhost:" + localstack.SQSPort),
		DisableSSL:  disableSSL,
	})
	if err != nil {
		t.Errorf("Could not create new AWS session: %v", err)
	}

	//when
	//Create SQS client cand call the list queues API
	sqsClient := sqs.New(sess)
	_, err = sqsClient.ListQueues(nil)

	//then
	//API did not produce an error
	assert.Nil(t, err, "ListQueues API should not produce an error")
}
