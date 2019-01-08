package localstack

import (
	"testing"

	"github.com/CallumKerrEdwards/go/container/receipt/assert"
)

func TestS3Session(t *testing.T) {
	//given
	asrt := assert.NewAssert(t)
	expectedEndpoint := "http://localhost:4572"

	//when
	sess := S3Session()

	//then
	actualEndpoint := *sess.Config.Endpoint
	asrt.Equal(actualEndpoint, expectedEndpoint, "Endpoints")
}

func TestSNSSession(t *testing.T) {
	//given
	asrt := assert.NewAssert(t)
	expectedEndpoint := "http://localhost:4575"

	//when
	sess := SNSSession()

	//then
	actualEndpoint := *sess.Config.Endpoint
	asrt.Equal(actualEndpoint, expectedEndpoint, "Endpoints")
}

func TestSQSSession(t *testing.T) {
	//given
	asrt := assert.NewAssert(t)
	expectedEndpoint := "http://localhost:4576"

	//when
	sess := SQSSession()

	//then
	actualEndpoint := *sess.Config.Endpoint
	asrt.Equal(actualEndpoint, expectedEndpoint, "Endpoints")
}
