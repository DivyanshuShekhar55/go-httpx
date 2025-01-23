package filehandler

import (
	"fmt"
	"os"
)

func CreateFile(fileName, text string) error {

	// create a new file ...
	file, err := os.Create("C:/my_stuff/go/go-httpx/content/Post" + fileName + ".txt")
	if err != nil {
		fmt.Println(err)
		return err
	}

	// write on the file ...
	// val returned on success is length but we don't need it so using an '_'

	_, err = file.WriteString(text)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}
