FROM golang:1.18

WORKDIR '/go-image-scraper'
COPY . .

RUN make

ENTRYPOINT ["./main"]
# SERVER_MODE = 1
CMD ["1"] 