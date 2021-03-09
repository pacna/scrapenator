package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Println("Enter url")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		url, err := url.ParseRequestURI(scanner.Text())
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

		document, err := goquery.NewDocumentFromReader(response.Body)

		if err != nil {
			log.Fatal("no html content")
		}

		document.Find("img").Each(func(index int, imgContent *goquery.Selection) {
			imgSrc, isSrcEmpty := imgContent.Attr("src")

			if isSrcEmpty {
				fmt.Println(imgSrc)
			}
		})
	}
}
