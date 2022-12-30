# Scrapenator

Scrapenator is a web scraper designed to extract images from web pages.

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
-   [Make](https://www.gnu.org/software/make/)
-   [Docker](https://docs.docker.com/get-docker/) (optional)
-   [Docker Compose](https://docs.docker.com/compose/install/) (optional)

### How to Run Locally

To run Scrapenator in terminal mode, use the following command:

```bash
$ make terminal
```

To run Scrapenator in server mode, execute the following command:

```bash
$ make server
```

### How to Run Using Docker (Optional)

To utilize Docker for Scrapenator, you can easily deploy it using the following command:

```bash
$ make docker
```

**Note:** Docker only runs in [SERVER MODE](#server-mode).
