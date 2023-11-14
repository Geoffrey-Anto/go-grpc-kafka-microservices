FROM golang:1.21

# Set destination for COPY
WORKDIR /logger-decoupled

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

EXPOSE ${PORT}

RUN go build -o ./logger-decoupled-exec

RUN chmod +x ./logger-decoupled-exec

CMD ./logger-decoupled-exec

