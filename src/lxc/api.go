package lxc

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// Connect docker server
func Connect(url string) (cli *client.Client, err error) {
	cli, err = client.NewClientWithOpts(client.WithHost(url), client.WithVersion("1.39"))
	return
}

// Run container
func Run(cli *client.Client, parseData []map[string]string) (runLog string, err error) {
	for _, el := range parseData {
		image := el["image"] + ":" + el["version"]
		runLogContainer, err := startContainer(cli, image, el["port"])
		if err != nil {
			break
		}
		runLog += runLogContainer
	}
	return
}

// start docker container
func startContainer(cli *client.Client, image, port string) (runLog string, err error) {
	ctx := context.Background()
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: image,
	}, nil, nil, "")

	if err != nil {
		return
	}

	if err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return
	}

	runLog = fmt.Sprintf("start container success, image: %s\n", image)
	log.Println(runLog)
	return
}
