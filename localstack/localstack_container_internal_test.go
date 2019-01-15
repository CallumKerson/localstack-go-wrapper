package localstack

import (
	"fmt"
	"testing"

	"github.com/callumkerredwards/localstack-go-wrapper/localstack/services"
	"github.com/stretchr/testify/assert"
)

func TestNoServicesProvidedToContainerConfig(t *testing.T) {
	//when
	container, err := New()

	assert.NoError(t, err, "New should not throw an error if docker is installed and working")
	assert.Len(t, container.containerConfig.Env, 0, "Should not provide any env variables if no services are set")
}

func TestServicesProvidedToContainerConfig(t *testing.T) {
	//given
	sqsConfig := &services.ServiceConfig{
		Service: services.SQS,
	}
	redshiftConfig := &services.ServiceConfig{
		Service: services.Redshift,
	}
	expectedEnv := "SERVICES=sqs,redshift"

	//when
	container, err := New(sqsConfig, redshiftConfig)

	//then
	assert.NoError(t, err, "New should not throw an error if docker is installed and working")
	assert.Len(t, container.containerConfig.Env, 1)
	assert.Equal(t, expectedEnv, container.containerConfig.Env[0])
}

func TestNoServicesProvidedToHostConfig(t *testing.T) {
	//when
	container, err := New()

	//then
	assert.NoError(t, err, "New should not throw an error if docker is installed and working")
	assert.Len(t, container.hostConfig.PortBindings, 17, "Should have mappings for all supported services")
	m := fmt.Sprintf("%v", container.hostConfig.PortBindings)
	assert.Contains(t, m, "4567/tcp:[{0.0.0.0 4567}]")
	assert.Contains(t, m, "4577/tcp:[{0.0.0.0 4577}]")
	assert.Contains(t, m, "4584/tcp:[{0.0.0.0 4584}]")
}

func TestServicesProvidedToHostConfig(t *testing.T) {
	//given
	sqsConfig := &services.ServiceConfig{
		Service: services.SQS,
		Port:    3003,
	}
	redshiftConfig := &services.ServiceConfig{
		Service: services.Redshift,
	}

	//when
	container, err := New(sqsConfig, redshiftConfig)

	//then
	assert.NoError(t, err, "New should not throw an error if docker is installed and working")
	assert.Len(t, container.hostConfig.PortBindings, 2)
	m := fmt.Sprintf("%v", container.hostConfig.PortBindings)
	assert.Contains(t, m, "4576/tcp:[{0.0.0.0 3003}]")
	assert.Contains(t, m, "4577/tcp:[{0.0.0.0 4577}]")
	assert.NotContains(t, m, "4567/tcp:[{0.0.0.0 4567}]")
	assert.NotContains(t, m, "4584/tcp:[{0.0.0.0 4584}]")
}
