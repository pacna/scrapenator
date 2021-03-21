package utils

import (
	"io"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Scrape(body io.Reader) []string {
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

	for imgUrl := range uniqueImgUrls {
		imgUrls = append(imgUrls, imgUrl)
	}

	return imgUrls
}

func UpdateUrl(userInput string) string {
	var updatedUrl string

	urlSegments := strings.Split(userInput, ".")

	// ex: google.com
	if len(urlSegments) == 2 {
		urlSegments[0] = "https://www." + urlSegments[0]
	}

	// ex: www.google.com
	if !strings.Contains(urlSegments[0], "http") {
		urlSegments[0] = "https://" + urlSegments[0]

	}

	updatedUrl = strings.Join(urlSegments, ".")

	return updatedUrl
}
