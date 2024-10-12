package main

import (
	"fmt"
	"sync"

	configs "github.com/princetanwar/file_downloader/internal/constants"
	utils "github.com/princetanwar/file_downloader/internal/utilities"
)

func main() {

	utils.CreateDownloadDir()

	ch := make(chan configs.FileConfig)
	var wg sync.WaitGroup

	for i := 0; i < configs.NumWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range ch {

				fmt.Printf("got file at worker %d, file name %s \n", i, file.Name)
				utils.DownloadFile(file.Url, file.Name)
			}
		}()
	}

	for i, f := range configs.Files {
		fmt.Println("adding at", i)
		ch <- f

	}
	close(ch)
	wg.Wait()

}
