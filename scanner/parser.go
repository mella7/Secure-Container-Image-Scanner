package scanner

import (
	"context"
	"encoding/json"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Layer struct {
	Digest string
}

// ParseImageLayers retrieves the layers of a Docker image
func ParseImageLayers(imageName string) ([]Layer, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	imageInspect, _, err := cli.ImageInspectWithRaw(context.Background(), imageName)
	if err != nil {
		return nil, err
	}

	layers := []Layer{}
	for _, layer := range imageInspect.RootFS.Layers {
		layers = append(layers, Layer{Digest: layer})
	}
	return layers, nil
}
