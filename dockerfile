FROM golang:1.12

WORKDIR /go/src/mutant-detector
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["mutant-detector"]