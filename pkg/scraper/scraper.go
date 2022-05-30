package scraper

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/asaskevich/govalidator"
)

func Scrape(updatedURL string, body io.Reader) []string {
	var uniqueImgURLs map[string]bool = make(map[string]bool)
	var imgURLs []string

	document, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		log.Fatal("no html content")
	}

	document.Find("img").Each(func(index int, imgContent *goquery.Selection) {
		imgSrc, isSrcEmpty := imgContent.Attr("src")
		dataImgSrc, isDataSrcEmpty := imgContent.Attr("data-src")

		if isSrcEmpty {
			uniqueImgURLs[imgSrc] = true
		}

		if isDataSrcEmpty {
			uniqueImgURLs[dataImgSrc] = true
		}
	})

	for imgURL := range uniqueImgURLs {
		if govalidator.IsURL(imgURL) {
			imgURLs = append(imgURLs, imgURL)
		}

		if !govalidator.IsURL(imgURL) {
			imgURLs = append(imgURLs, updatedURL+imgURL)
		}
	}

	return imgURLs
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

func DownloadImages(imgURLs []string) error {
	zipFile, _ := os.Create(strconv.FormatInt(time.Now().Unix(), 10) + ".zip")
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for index, imgURL := range imgURLs {
		var imgURLInSegments []string = strings.Split(imgURL, "/")
		var fileName string = createFileName(imgURLInSegments[len(imgURLInSegments)-1], index)

		var zipInfoReceiver ZipInfoReceiver
		zipInfoReceiver.fileName = fileName
		zipInfoReceiver.imgURL = imgURL
		zipInfoReceiver.zipWriter = zipWriter

		err := appendImageToZip(zipInfoReceiver)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetUpdatedURL(inputURL string) string {
	if !govalidator.IsURL(inputURL) {
		return ""
	}

	response, err := http.Get(inputURL)

	if err != nil {
		return ""
	}

	defer response.Body.Close()

	updatedURL := fmt.Sprintf("%s://%s", response.Request.URL.Scheme, response.Request.URL.Host)

	return updatedURL;
}

func GetResponseFromURL(inputURL string) io.Reader {
	var buffer bytes.Buffer
	response, err := http.Get(inputURL)

	if err != nil {
		return nil
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil
	}

	buffer.ReadFrom(response.Body)

	responseBody := ioutil.NopCloser(&buffer)

	return responseBody
}

func createFileName(fileNameFromURL string, index int) string {
	reg, _ := regexp.Compile(`[\d?]+`)
	fileNameFromURLSegments := strings.Split(fileNameFromURL, ".")

	cleanFileName := reg.ReplaceAllString(fileNameFromURLSegments[0], "")
	cleanFileExtension := reg.ReplaceAllString(fileNameFromURLSegments[1], "")

	fileName := fmt.Sprintf("%s_%d.%s", cleanFileName, index, cleanFileExtension)

	return fileName
}

func appendImageToZip(zipInfo ZipInfoGetter) error {
	image := storeImage(zipInfo.getImgUrl())
	zipFileHeader := &zip.FileHeader{
		Name:   zipInfo.getFileName(),
		Method: zip.Deflate,
	}

	zipFile, _ := zipInfo.getZipWriter().CreateHeader(zipFileHeader)

	_, err := io.Copy(zipFile, image)

	if err != nil {
		return err
	}

	return nil
}
