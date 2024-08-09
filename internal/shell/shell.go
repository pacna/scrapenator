package shell

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pacna/goober/internal/scraper"
	"github.com/pacna/goober/internal/utility"
)

func New() {
	processUserInput()
}

func processUserInput() {
	fmt.Print("Enter url ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inputURL := scanner.Text()

		if inputURL == "q" {
			fmt.Println("Bye | (• ◡•)| (❍ᴥ❍ʋ)")
			break
		}

		if inputURL != "q" {

			updatedURL := scraper.GetUpdatedURL(inputURL)
			responseBody := scraper.GetResponseFromURL(inputURL)
			
			if (utility.IsStringEmpty(updatedURL) || responseBody == nil) {
				fmt.Println("Unable to process URL")
				break;
			}

			imgURLs := scraper.Scrape(updatedURL, responseBody)
			scraper.DownloadImages(imgURLs)
			fmt.Print("Enter another url or press q to QUIT ")
		}
	}
}