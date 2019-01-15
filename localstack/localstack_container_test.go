package localstack

import (
	"context"
	"testing"
	"unicode/utf8"

	"github.com/callumkerredwards/localstack-go-wrapper/localstack/services"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/stretchr/testify/assert"
)

func TestCreatesContainer(t *testing.T) {
	//when
	container, err := New()

	//then
	assert.NoError(t, err, "New should not throw an error if docker is installed and working")
	assert.NotNil(t, container.ContainerID)
	assert.True(t, utf8.RuneCountInString(container.ContainerID) > 0)
	assert.NotContains(t, runningDockerContainerIds(t), container.ContainerID, "New should not start container")
}

func TestStartsContainer(t *testing.T) {
	//given
	container, err := New()
	assert.NoError(t, err)

	//when
	err = container.Start()
	defer container.Stop()

	//then
	assert.NoError(t, err, "Start should start docker container successfully")
	assert.Contains(t, runningDockerContainerIds(t), container.ContainerID)
}

func TestStopsContainer(t *testing.T) {
	//given
	container, err := New()
	assert.NoError(t, err)
	err = container.Start()
	assert.NoError(t, err)

	//when
	err = container.Stop()

	//then
	assert.NoError(t, err, "Stop should stop docker container successfully")
	assert.NotContains(t, runningDockerContainerIds(t), container.ContainerID)
}

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

func runningDockerContainerIds(t *testing.T) []string {
	cli, err := client.NewEnvClient()
	if err != nil {
		t.Error(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		t.Error(err)
	}

	ids := make([]string, 0, len(containers))
	for _, c := range containers {
		ids = append(ids, c.ID)
	}
	return ids
}
