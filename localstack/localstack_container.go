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

const ImageName = "localstack/localstack"

var dockerClient *client.Client = func() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return cli
}

func Start() string {
	ctx := context.Background()
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	normalized, err := reference.ParseNormalizedNamed(ImageName)
	if err != nil {
		panic(err)
	}

	out, err := dockerClient.ImagePull(ctx, normalized.String(), types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)

	containerConfig := &container.Config{
		Image: ImageName,
		Env:   []string{"SERVICES=s3,sns,sqs"},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"4572/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "4572",
				},
			},
			"4575/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "4575",
				},
			},
			"4576/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "4576",
				},
			},
		},
	}

	resp, err := dockerClient.ContainerCreate(ctx, containerConfig, hostConfig, nil, "")
	if err != nil {
		panic(err)
	}

	if err := dockerClient.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	duration := time.Second * 5
	time.Sleep(duration)
	return resp.ID
}

func Stop(id string) {
	ctx := context.Background()
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	if err := dockerClient.ContainerStop(ctx, id, nil); err != nil {
		panic(err)
	}
}
