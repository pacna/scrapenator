FROM golang:1.21

WORKDIR '/scrapenator'
COPY . .

RUN make build

# Set the entry point to the first binary file found in the current directory
ENTRYPOINT ["./scrapenator"]