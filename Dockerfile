FROM golang:1.22.5-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main cmd/main.go

CMD ["/app/main"]
