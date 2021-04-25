package main

import (
	"bufio"
	"fmt"
	"go-image-scraper/utils"
	"log"
	"net/http"
	"os"

	"github.com/asaskevich/govalidator"
)

func main() {
	fmt.Print("Enter url ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputURL := scanner.Text()

		if !govalidator.IsURL(inputURL) {
			log.Fatal("Invalid url", inputURL)
		}

		response, err := http.Get(inputURL)

		updatedURL := "http://" + response.Request.URL.Host

		if err != nil {
			log.Fatal("Invalid request")
		}

		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			log.Fatal("Status is not returning a success code", response.StatusCode, response.Status)
		}

		imgUrls := utils.Scrape(updatedURL, response.Body)

		// utils.DownloadImages(imgUrls)

		fmt.Println(imgUrls)
	}
}
