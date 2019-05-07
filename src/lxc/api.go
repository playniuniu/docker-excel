package lxc

import (
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
		runLogContainer, err := startContainer(cli, el)
		if err != nil {
			break
		}
		runLog += runLogContainer
	}
	return
}
