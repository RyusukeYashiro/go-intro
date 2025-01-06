FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go mod download && go build -o main .