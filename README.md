# Scrapenator

Scrapenator is a web scraper designed to extract images from web pages. It provides both terminal and server modes to download images or return a list of images in the response.

### TERMINAL MODE

In terminal mode, Scrapenator downloads images and packs them into a zip file.

```bash
Enter url https://www.google.com/
Enter another url or press q to QUIT q
Bye | (• ◡•)| (❍ᴥ❍ʋ)
```

### SERVER MODE

In server mode, Scrapenator acts as a basic service that returns a list of images in the response. It runs locally at http://localhost:5000/.

```http
http://localhost:5000/scrape?url=https://www.google.com/
```

```json
{
    "imgs": [
        "https://www.google.com/images/branding/googlelogo/1x/googlelogo_white_background_color_272x92dp.png"
    ]
}
```

## Prerequisites

Before using Scrapenator, make sure you have the following tools and components installed:

-   [Golang](https://golang.org/dl/)
-   [Docker](https://docs.docker.com/get-docker/) (optional)
-   [Docker Compose](https://docs.docker.com/compose/install/) (optional)

### How to Run Locally

To run in terminal mode:

```bash
$ make terminal
```

To run in server mode:

```bash
$ make server
```

### How to Run Using Docker (Optional)

Navigate to the deployments directory:

```bash
$ cd deployments
```

Run the following command:

```bash
$ docker-compose up --build
```

Or you can use the Make command:

```bash
$ make docker
```

**Note:** Docker only runs in [SERVER MODE](#server-mode).
