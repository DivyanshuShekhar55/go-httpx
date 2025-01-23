package filehandler

import (
	"os"
)

func GetFile(filePath string) ([]byte, error) {
	file_fullPath := "C:/my_stuff/go/go-httpx/content/" + filePath + ".txt"

	data, err := os.ReadFile(file_fullPath)
	if err != nil {
		return nil, err
	}

	return data, nil
}
