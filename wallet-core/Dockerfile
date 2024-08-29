FROM golang:1.23 AS build-stage

WORKDIR /go/src/app

RUN apt-get update && apt-get install -y librdkafka-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# RUN go build -o app ./cmd/main.go

CMD ["go", "run", "./cmd/main.go"]
