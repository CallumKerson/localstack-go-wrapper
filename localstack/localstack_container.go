package localstack

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/docker/distribution/reference"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// ImageName provides the docker image name for LocalStack
const ImageName = "localstack/localstack"

// S3Port provides the port that the LocalStack S3 service is running on
const S3Port = "4572"

// SNSPort provides the port that the LocalStack SNS service is running on
const SNSPort = "4575"

// SQSPort provides the port that the LocalStack SQS service is running on
const SQSPort = "4576"

// Start pulls, creates and starts a LocalStack container then returning the id of the container
func Start() (string, error) {
	ctx := context.Background()
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		return "nil", err
	}

	normalized, err := reference.ParseNormalizedNamed(ImageName)
	if err != nil {
		return "nil", err
	}

	out, err := dockerClient.ImagePull(ctx, normalized.String(), types.ImagePullOptions{})
	if err != nil {
		return "nil", err
	}
	if _, err = io.Copy(os.Stdout, out); err != nil {
		return "nil", err
	}

	containerConfig := &container.Config{
		Image: ImageName,
		Env:   []string{"SERVICES=s3,sns,sqs"},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"4572/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: S3Port,
				},
			},
			"4575/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: SNSPort,
				},
			},
			"4576/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: SQSPort,
				},
			},
		},
	}

	resp, err := dockerClient.ContainerCreate(ctx, containerConfig, hostConfig, nil, "")
	if err != nil {
		return "nil", err
	}

	if err := dockerClient.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "nil", err
	}

	duration := time.Second * 5
	time.Sleep(duration)
	return resp.ID, nil
}

// Stop stop the docker container identified by the id string
func Stop(id string) error {
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		return err
	}
	if err := dockerClient.ContainerStop(context.Background(), id, nil); err != nil {
		return err
	}
	return nil
}
