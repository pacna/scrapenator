package main

import (
	"bufio"
	"fmt"
	"go-image-extractor/utils"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	fmt.Print("Enter url ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		updatedUrl := utils.UpdateUrl(scanner.Text())

		url, err := url.ParseRequestURI(updatedUrl)

		if err != nil {
			log.Fatal("Invalid url")
		}

		response, err := http.Get(url.String())

		if err != nil {
			log.Fatal("Invalid request")
		}

		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			log.Fatal("Status is not returning a success code", response.StatusCode, response.Status)
		}

		imgUrls := utils.Scrape(response.Body)

		fmt.Println(imgUrls)
	}
}
