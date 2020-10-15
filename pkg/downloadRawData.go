package pkg

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(filepath string, url string) error {

	//Get data about url
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//Create file to save data in
	out, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	//Write data to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
