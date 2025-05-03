FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
RUN go build -o migrate ./database/cmd/migrate


FROM builder
EXPOSE 5000
CMD "./main"
