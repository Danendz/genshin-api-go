FROM golang:1.22-alpine3.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN GOOS=linux go build -o ./bin/server

CMD ["./bin/server"]