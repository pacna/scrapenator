FROM golang:1.19

WORKDIR '/scrapenator'
COPY . .

RUN make build-server

ENTRYPOINT ["./server"]