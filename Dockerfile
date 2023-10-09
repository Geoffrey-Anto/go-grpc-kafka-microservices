FROM golang:1.21

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN mkdir bin

ENV PORT 8080

RUN go build -o ./bin/server

CMD ["./bin/server"]
