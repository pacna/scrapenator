package utils

import (
	"archive/zip"
	"bytes"
	"go-image-scraper/utils/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

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

// DownloadImages -- download list of images
func DownloadImages(imgURLs []string) {
	zipFile, _ := os.Create(strconv.FormatInt(time.Now().Unix(), 10) + ".zip")
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, imgURL := range imgURLs {
		var imgURLInSegments []string = strings.Split(imgURL, "/")
		var fileName string = imgURLInSegments[len(imgURLInSegments)-1]

		var zipInfo models.ZipInfo
		zipInfo.ZipFile = zipFile
		zipInfo.ZipWriter = zipWriter
		zipInfo.FileName = fileName
		zipInfo.ImgURL = imgURL

		appendImageToZip(zipInfo)
	}
}

func appendImageToZip(zipInfo models.ZipInfo) error {
	image := storeImage(zipInfo.ImgURL)
	zipFileHeader := &zip.FileHeader{
		Name:   zipInfo.FileName,
		Method: zip.Deflate,
	}

	zipFile, _ := zipInfo.ZipWriter.CreateHeader(zipFileHeader)

	_, err := io.Copy(zipFile, image)

	if err != nil {
		return err
	}

	return nil
}
