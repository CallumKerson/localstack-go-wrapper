package services_test

import (
	"testing"

	"github.com/callumkerredwards/localstack-go-wrapper/localstack/services"
	"github.com/stretchr/testify/assert"
)

func TestSupportedServices(t *testing.T) {
	assert.Len(t, services.SupportedServices, 17)
}
