FROM golang:latest

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download && go mod verify

COPY . .
EXPOSE 8080
RUN go build
CMD ["./AnekaZoo"]
