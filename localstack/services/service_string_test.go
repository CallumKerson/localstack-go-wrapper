package services_test

import (
	"testing"

	"github.com/callumkerredwards/localstack-go-wrapper/localstack/services"
	"github.com/stretchr/testify/assert"
)

func TestAPIGatewayString(t *testing.T) {
	assert.Equal(t, "APIGateway", services.APIGateway.String())
}

func TestS3String(t *testing.T) {
	assert.Equal(t, "S3", services.S3.String())
}
