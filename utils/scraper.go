package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/asaskevich/govalidator"
)

// Scrape -- scapes img urls from html document
func Scrape(updatedURL string, body io.Reader) []string {
	var uniqueImgUrls map[string]bool = make(map[string]bool)
	var imgUrls []string

	document, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		log.Fatal("no html content")
	}

	document.Find("img").Each(func(index int, imgContent *goquery.Selection) {
		imgSrc, isSrcEmpty := imgContent.Attr("src")
		dataImgSrc, isDataSrcEmpty := imgContent.Attr("data-src")

		if isSrcEmpty {
			uniqueImgUrls[imgSrc] = true
		}

		if isDataSrcEmpty {
			uniqueImgUrls[dataImgSrc] = true
		}
	})

	for imgURL := range uniqueImgUrls {
		if govalidator.IsURL(imgURL) {
			imgUrls = append(imgUrls, imgURL)
		} else {
			imgUrls = append(imgUrls, updatedURL+imgURL)
		}
	}

	return imgUrls
}

// UpdateURL -- updates the user input if it does not have the minimum requirement of a url
func UpdateURL(userInput string) string {
	var updatedURL string

	urlSegments := strings.Split(userInput, ".")

	// ex: google.com
	if len(urlSegments) == 2 {
		urlSegments[0] = "https://www." + urlSegments[0]
	}

	// ex: www.google.com
	if !strings.Contains(urlSegments[0], "http") {
		urlSegments[0] = "https://" + urlSegments[0]

	}

	updatedURL = strings.Join(urlSegments, ".")

	return updatedURL
}

func storeImage(imgURL string) io.Reader {
	var buffer bytes.Buffer
	response, err := http.Get(imgURL)

	if err != nil {
		log.Fatal("invalid url")
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatal("Status is not returning a success code", response.StatusCode, response.Status)
	}

	buffer.ReadFrom(response.Body)

	imageBody := ioutil.NopCloser(&buffer)

	return imageBody
}

// CreateImage -- creates an image
func CreateImage(imgUrls []string) {
	image := storeImage(imgUrls[0])

	file, err := os.Create("temp.png")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, err = io.Copy(file, image)

	if err != nil {
		log.Fatal("could create image", err)
	}
}
