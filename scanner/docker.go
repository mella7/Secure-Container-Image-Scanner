package scanner

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// PullDockerImage pulls the specified Docker image
func PullDockerImage(imageName string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	defer cli.Close()

	_, err = cli.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
	return err
}
