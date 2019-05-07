package web

import (
	"log"
	"lxc"

	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

var cli *client.Client
var router *gin.Engine

func initDocker(URL string) {
	var err error
	cli, err = lxc.Connect(URL)
	if err != nil {
		log.Fatalf("failed to connect docker: %s\n", err)
	}
	log.Printf("Connect to docker: %s\n", URL)
}

func initGinServer(staticPath string) {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	router.Static("/", staticPath)
	router.POST("/upload", handleUpload)
}
