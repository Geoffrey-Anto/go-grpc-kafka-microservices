FROM golang:1.21

# Set destination for COPY
WORKDIR /server

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

EXPOSE ${PORT}

RUN go build -o ./logger-exec

RUN chmod +x ./logger-exec

CMD ./logger-exec --port $PORT

