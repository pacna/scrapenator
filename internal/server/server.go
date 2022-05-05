package server

import (
	"encoding/json"
	"fmt"
	"go-image-scraper/pkg/scraper"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

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

	var imgsResponse ImgURLResponse
	if len(queryValue) > 0 {
			
		updatedURL := scraper.GetUpdatedURL(queryValue)
		responseBody := scraper.GetResponseFromURL(queryValue)
		
		if (len(updatedURL) == 0 || responseBody == nil) {
			panic("Unable to process URL")

		}

		if responseBody != nil {
			var imgs []string = scraper.Scrape(updatedURL, responseBody)
			imgsResponse.Imgs = imgs
		}

		if responseBody == nil {
			emptyResponse := []string{}
			imgsResponse.Imgs = emptyResponse
		}
	}

	if len(queryValue) == 0 {
		emptyResponse := []string{}
		imgsResponse.Imgs = emptyResponse
	}

	writer.Header().Set("Content-Type", "application/json")

	response, _ := json.Marshal(imgsResponse)
	writer.Write(response)
}