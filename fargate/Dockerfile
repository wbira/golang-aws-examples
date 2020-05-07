FROM golang:1.13.0-alpine

WORKDIR /app

COPY go.mod .
#COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./src/server ./src

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./src/server"]