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
		updatedURL := utils.UpdateURL(scanner.Text())

		if !govalidator.IsURL(updatedURL) {
			log.Fatal("Invalid url", updatedURL)
		}

		response, err := http.Get(updatedURL)

		if err != nil {
			log.Fatal("Invalid request")
		}

		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			log.Fatal("Status is not returning a success code", response.StatusCode, response.Status)
		}

		imgUrls := utils.Scrape(updatedURL, response.Body)

		// utils.CreateImage(imgUrls)

		fmt.Println(imgUrls)
	}
}
