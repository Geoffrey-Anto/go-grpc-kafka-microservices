FROM golang:1.21

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

ENV PORT 3002

RUN go build -o ./server-exec

CMD ["./server-exec"]
