FROM golang:1.19

WORKDIR '/goscrape'
COPY . .

RUN make build-server

ENTRYPOINT ["./server"]