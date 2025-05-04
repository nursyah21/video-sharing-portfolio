FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

FROM builder
EXPOSE 5000
CMD ["./main"]
