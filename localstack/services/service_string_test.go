package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIGatewayString(t *testing.T) {
	assert.Equal(t, "APIGateway", APIGateway.String())
}

func TestS3String(t *testing.T) {
	assert.Equal(t, "S3", S3.String())
}
