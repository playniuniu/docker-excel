package excel

import (
	"net/http"
	"os"
)

// DetectFile detect file type
func DetectFile(filename string) (fileType string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	buff := make([]byte, 512)
	_, err = f.Read(buff)
	if err != nil {
		return
	}

	fileType = http.DetectContentType(buff)
	return
}
