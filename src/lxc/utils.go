package lxc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// parse host config, only support one port pair
// REFERENCE: https://medium.com/backendarmy/controlling-the-docker-engine-in-go-d25fc0fe2c45
func parseHostConfig(data map[string]string) (hostConfig *container.HostConfig, err error) {
	hostConfig = &container.HostConfig{}

	// config auto remove
	if data["remove"] == "true" {
		hostConfig.AutoRemove = true
	}

	// config port forward
	if data["port"] != "" {
		portPair := strings.Split(data["port"], ":")
		if len(portPair) != 2 {
			err = errors.New("port config error")
			return
		}

		hostBinding := nat.PortBinding{
			HostIP:   "0.0.0.0",
			HostPort: portPair[0],
		}
		containerPort, err := nat.NewPort("tcp", portPair[1])
		if err != nil {
			panic(err)
		}
		portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
		hostConfig.PortBindings = portBinding
	}

	return
}

// parse image name
func parseImageName(data map[string]string) string {
	if data["version"] != "" {
		return data["image"] + ":" + data["version"]
	}
	return data["image"]
}

// start docker container
func startContainer(cli *client.Client, data map[string]string) (runLog string, err error) {
	ctx := context.Background()

	imageName := parseImageName(data)
	config := &container.Config{
		Image: imageName,
		Env:   parseDockerEnv(data),
	}

	hostConfig, err := parseHostConfig(data)
	if err != nil {
		return
	}

	resp, err := cli.ContainerCreate(ctx, config, hostConfig, nil, data["name"])
	if err != nil {
		return
	}

	if err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return
	}

	runLog = fmt.Sprintf("start container success, image -> %s\n", imageName)
	log.Print(runLog)
	return
}

// parse docker env
func parseDockerEnv(data map[string]string) (env []string) {
	if data["env"] != "" {
		envList := strings.Split(data["env"], "&")
		for _, el := range envList {
			env = append(env, el)
		}
	}
	return
}
