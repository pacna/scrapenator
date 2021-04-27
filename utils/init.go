package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"go-image-scraper/utils/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
)

// InitTerminalMode -- initiate terminal mode
func InitTerminalMode() {
	processUserInput()
}

// InitServerMode -- initiate server mode
func InitServerMode() {
	router := setupRouter()

	server := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:5000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server is listening on PORT 5000")
	log.Fatal(server.ListenAndServe())
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/scraper", scraperHandler)
	return router
}

func scraperHandler(writer http.ResponseWriter, request *http.Request) {
	queryValue := request.URL.Query().Get("url")

	var imgsResponse models.ImgURLResponse
	if len(queryValue) > 0 {
		updatedURL, responseBody := processURL(queryValue)

		if responseBody != nil {
			var imgs []string = scrape(updatedURL, responseBody)
			imgsResponse.Imgs = imgs
		} else {
			emptyResponse := []string{}
			imgsResponse.Imgs = emptyResponse
		}
	} else {
		emptyResponse := []string{}
		imgsResponse.Imgs = emptyResponse
	}
	writer.Header().Set("Content-Type", "application/json")

	response, _ := json.Marshal(imgsResponse)
	writer.Write(response)
}

func processURL(inputURL string) (string, io.Reader) {
	var buffer bytes.Buffer
	if !govalidator.IsURL(inputURL) {
		return "", nil
	}

	response, err := http.Get(inputURL)
	updatedURL := fmt.Sprintf("%s://%s", response.Request.URL.Scheme, response.Request.URL.Host)

	if err != nil {
		return "", nil
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", nil
	}

	buffer.ReadFrom(response.Body)

	responseBody := ioutil.NopCloser(&buffer)

	return updatedURL, responseBody
}

func processUserInput() {
	fmt.Print("Enter url ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inputURL := scanner.Text()

		if inputURL == "q" {
			fmt.Println("Bye | (• ◡•)| (❍ᴥ❍ʋ)")
			break
		} else {
			updatedURL, responseBody := processURL(inputURL)

			if responseBody == nil {
				fmt.Println("Something went wrong")
				break
			}

			imgURLs := scrape(updatedURL, responseBody)
			downloadImages(imgURLs)
			fmt.Print("Enter another url or press q to QUIT ")
		}

	}
}
