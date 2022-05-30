FROM golang:1.18

WORKDIR '/go-image-scraper'
COPY . .

RUN make build

ENTRYPOINT ["./server"]