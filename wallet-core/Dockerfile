FROM golang:1.23 AS build-stage

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/main.go

CMD ["./app"]
