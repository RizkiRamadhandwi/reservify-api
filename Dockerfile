FROM golang:alpine

WORKDIR /app

COPY . .
COPY .env /app

RUN go mod tidy
RUN go build -o reservify-app

ENTRYPOINT ["/app/reservify-app"]
