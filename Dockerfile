FROM golang:1.14

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v cmd/main.go

CMD ["main"]

EXPOSE 3000