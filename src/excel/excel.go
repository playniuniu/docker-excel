package excel

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

// ReadFile read excel file
func ReadFile(filename string) (res []map[string]string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	// detect file type
	var fileExt string
	if _, fileExt, err = mimetype.DetectReader(f); err != nil {
		return
	}
	// Important! need to reset file point
	f.Seek(0, io.SeekStart)

	// read excel data
	res = []map[string]string{}
	switch fileExt {
	case "csv":
		res = parseCSVData(f)
	case "zip":
		res = parseExcelData(f)
	default:
		err = errors.New("unsupport file type")
	}
	return
}

func parseCSVData(reader io.Reader) []map[string]string {
	r := csv.NewReader(reader)
	rows := []map[string]string{}
	var header []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if header == nil {
			header = record
		} else {
			dict := map[string]string{}
			for i := range header {
				dict[header[i]] = record[i]
			}
			rows = append(rows, dict)
		}
	}
	return rows
}

func parseExcelData(reader io.Reader) []map[string]string {
	return nil
}
