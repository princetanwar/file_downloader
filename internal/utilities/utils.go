package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"

	constants "github.com/princetanwar/file_downloader/internal/constants"
)

func storeFile(data []byte, fileName string) {

	filepath := fmt.Sprintf("%s/%s", constants.DOWNLOAD_DIR, fileName)
	file, err := os.Create(filepath)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	write, err := file.Write(data)

	if err != nil {
		panic(err)
	}
	fmt.Printf("able to write %d \n", write)

}

func CreateDownloadDir() {

	err := os.Mkdir(constants.DOWNLOAD_DIR, 0755)

	isAlreadyExist := os.IsExist(err)

	if isAlreadyExist {
		fmt.Println("directory already exist")
		return
	}

	if err != nil {
		fmt.Println("failed to create directory", err)
	} else {
		fmt.Println("successfully created directory")
	}
}

func DownloadFile(url string, saveFileName string) {

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("got error", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("error while reading from body", err)
	}

	storeFile(body, saveFileName)

}
