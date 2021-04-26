package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
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
	router.HandleFunc("/scraper", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		response, _ := json.Marshal(struct{ Message string }{
			Message: "Hello World",
		})
		writer.Write(response)
	})
	return router
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
			if !govalidator.IsURL(inputURL) {
				fmt.Println("Invalid URL ------", inputURL)
				break
			}

			response, err := http.Get(inputURL)
			updatedURL := fmt.Sprintf("%s://%s", response.Request.URL.Scheme, response.Request.URL.Host)

			if err != nil {
				fmt.Println(err)
				break
			}

			defer response.Body.Close()

			if response.StatusCode != http.StatusOK {
				fmt.Println("Status is not returning a success code", response.StatusCode, response.Status)
				break
			}

			imgURLs := scrape(updatedURL, response.Body)

			// downloadImages(imgURLs)

			fmt.Println(imgURLs)
			fmt.Print("Enter another url or press q to QUIT ")
		}

	}
}
