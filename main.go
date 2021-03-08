package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	fmt.Println("Enter url")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		url, err := url.ParseRequestURI(scanner.Text())
		if err != nil {
			panic(err)
		}

		response, err := http.Get(url.String())

		if err != nil {
			panic(err)
		}

		defer response.Body.Close()
		var content []byte

		if content, err = ioutil.ReadAll(response.Body); err != nil {
			panic(err)
		}

		fmt.Println("Url", string(content))
	}
}
