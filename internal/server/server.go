package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pacna/goscrape/pkg/scraper"
	"github.com/pacna/goscrape/pkg/utility"

	"github.com/gorilla/mux"
)

// ImgURLCollection -- response model for returning a list of images from a webpage
type ImgURLCollection struct {
	Imgs []string `json:"imgs"`
}

// New -- initiate server mode
func New() {
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
	router.HandleFunc("/scrape", scrapeHandler)
	return router
}

func scrapeHandler(writer http.ResponseWriter, request *http.Request) {
	queryValue := request.URL.Query().Get("url")

	var imgCollection ImgURLCollection
	if !utility.IsStringEmpty(queryValue) {
			
		updatedURL := scraper.GetUpdatedURL(queryValue)
		responseBody := scraper.GetResponseFromURL(queryValue)
		
		if (utility.IsStringEmpty(updatedURL) || responseBody == nil) {
			fmt.Println("Unable to process URL")
			return
		}

		if responseBody != nil {
			var imgs []string = scraper.Scrape(updatedURL, responseBody)
			imgCollection.Imgs = imgs
		}

		if responseBody == nil {
			emptyResponse := []string{}
			imgCollection.Imgs = emptyResponse
		}
	}

	if utility.IsStringEmpty(queryValue) {
		emptyResponse := []string{}
		imgCollection.Imgs = emptyResponse
	}

	writer.Header().Set("Content-Type", "application/json")

	response, _ := json.Marshal(imgCollection)
	writer.Write(response)
}