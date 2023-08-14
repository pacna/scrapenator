FROM golang:1.21

WORKDIR '/scrapenator'
COPY . .

RUN make build_server

ENTRYPOINT ["./server"]