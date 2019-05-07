package web

import (
	"excel"
	"fmt"
	"log"
	"lxc"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Run web server
func Run(staticPath, port, dockerURL string) {
	initDocker(dockerURL)
	initGinServer(staticPath)
	log.Printf("Listen on port: %s\n", port)
	router.Run(port)
}

func handleUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Base(file.Filename)
	filename = filepath.Join("upload", filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	parseData := []map[string]string{}
	if parseData, err = excel.ReadFile(filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("read file err: %s", err.Error()))
		return
	}

	var runLog string
	if runLog, err = lxc.Run(cli, parseData); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("run docker err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("Run %s success\n%s", file.Filename, runLog))
}
