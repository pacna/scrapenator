# Goscrape

A web scraper that scrape images off of a web page.

### TERMINAL MODE

Download the images in a zip file.

```bash
Enter url https://www.google.com/
Enter another url or press q to QUIT q
Bye | (• ◡•)| (❍ᴥ❍ʋ)
```

### SERVER MODE

A basic service that returns a list of images in the response. Runs locally on http://localhost:5000/.

```
http://localhost:5000/scraper?url=https://www.google.com/

{"imgs":["https://www.google.com/images/branding/googlelogo/1x/googlelogo_white_background_color_272x92dp.png"]}
```

### Prerequisites

-   [Golang](https://golang.org/dl/)
-   [Docker](https://docs.docker.com/get-docker/) (optional)
-   [Docker Compose](https://docs.docker.com/compose/install/) (optional)

### How to run locally

```bash
# terminal mode
$ make terminal

# server mode
$ make server
```

### How to run using docker (Optional)

```bash
# goto the deployments directory
$ cd deployments

# run cmd
$ docker-compose up --build

# or use Make cmd
$ make docker
```

##### note -- docker only runs in SERVER MODE
