package integration

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
)

func TestLocalstackS3Started(t *testing.T) {
	//given
	//AWS session
	sess, err := session.NewSession(&aws.Config{
		Credentials: testCredentials,
		Region:      testRegion,
		Endpoint:    aws.String("http://localhost:" + strconv.Itoa(s3Port)), // uses custom s3 port
		DisableSSL:  disableSSL,
	})
	if err != nil {
		t.Errorf("Could not create new AWS session: %v", err)
	}
	s, err := json.MarshalIndent(sess, "", "  ")
	if err == nil {
		t.Logf("AWS Session is %s", string(s))
	}

	//when
	//Create S3 service client and call the list bucket API
	s3Client := s3.New(sess)
	_, err = s3Client.ListBuckets(nil)

	//then
	//API did not produce an error
	assert.Nil(t, err, "ListBuckets API should not produce an error")
}
