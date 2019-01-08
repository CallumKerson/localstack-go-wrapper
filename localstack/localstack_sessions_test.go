package localstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestS3Session(t *testing.T) {
	//given
	expectedEndpoint := "http://localhost:4572"

	//when
	sess := S3Session()

	//then
	actualEndpoint := *sess.Config.Endpoint
	assert.Equal(t, expectedEndpoint, actualEndpoint,
		"session endpoint should be localstack default for S3")
}

func TestSNSSession(t *testing.T) {
	//given
	expectedEndpoint := "http://localhost:4575"

	//when
	sess := SNSSession()

	//then
	actualEndpoint := *sess.Config.Endpoint
	assert.Equal(t, expectedEndpoint, actualEndpoint,
		"session endpoint should be localstack default for SNS")
}

func TestSQSSession(t *testing.T) {
	//given
	expectedEndpoint := "http://localhost:4576"

	//when
	sess := SQSSession()

	//then
	actualEndpoint := *sess.Config.Endpoint
	assert.Equal(t, expectedEndpoint, actualEndpoint,
		"session endpoint should be localstack default for SQS")
}
