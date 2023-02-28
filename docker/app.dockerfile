FROM golang:latest AS development

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go install github.com/mitranim/gow@latest

EXPOSE 3000

CMD gow run cmd/main/main.go
