package main

import (
	"web"
)

func main() {
	dockerURL := "unix:///var/run/docker.sock"
	web.Run("./web", ":9001", dockerURL)
}
