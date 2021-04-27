# GO Image Scraper

Scrape images off of a web page.

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

- [golang](https://golang.org/dl/)
- [docker](https://docs.docker.com/get-docker/) (optional)
- [docker-compose](https://docs.docker.com/compose/install/) (optional)

### How to run

```bash
$ make
# defaults to TERMINAL MODE
$ ./main
# passing in 1 will enable SERVER MODE
$ ./main 1
```

### How to run using docker

```bash
$ docker-compose up --build
```

##### note -- docker only runs in SERVER MODE
