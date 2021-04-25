package utils

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"github.com/asaskevich/govalidator"
)

// InitTerminalMode -- initiate terminal mode
func InitTerminalMode() {
	processUserInput()
}

// InitServerMode -- initiate server mode
func InitServerMode() {
	fmt.Println("Listening on port 5000")
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
