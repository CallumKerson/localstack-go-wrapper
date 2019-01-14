package integration

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/callumkerredwards/localstack-go-wrapper/localstack"
	"github.com/stretchr/testify/assert"
)

func TestLocalstackS3Started(t *testing.T) {
	//given
	//AWS session
	sess, err := session.NewSession(&aws.Config{
		Credentials: testCredentials,
		Region:      testRegion,
		Endpoint:    aws.String("http://localhost:" + localstack.S3Port),
		DisableSSL:  disableSSL,
	})
	if err != nil {
		t.Errorf("Could not create new AWS session: %v", err)
	}

	//when
	//Create S3 service client and call the list bucket API
	s3Client := s3.New(sess)
	_, err = s3Client.ListBuckets(nil)

	//then
	//API did not produce an error
	assert.Nil(t, err, "ListBuckets API should not produce an error")
}
