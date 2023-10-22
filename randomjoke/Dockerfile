FROM golang:1.21

# Set destination for COPY
WORKDIR /randomjoke

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

EXPOSE ${PORT}

RUN go build -o ./randomjoke-exec

RUN chmod +x ./randomjoke-exec

CMD ./randomjoke-exec --port $PORT

