package main

import (
	"web"
)

func main() {
	dockerURL := "tcp://47.75.200.118:23750"
	web.Run("./web", ":9001", dockerURL)
}
