#build stage
FROM golang:1.19.6-alpine3.16 AS builder
RUN apk add --no-cache git
WORKDIR /app

# download dependencies
COPY go.* .

RUN go mod download
RUN go mod verify

# build app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pokedex ./cmd/webserver

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates


WORKDIR /app

COPY --from=builder /app/pokedex .

# inject .env file, may not be the best practice, but will do for now
COPY .env .

RUN ls -l

ENTRYPOINT ["./pokedex", "--config", ".env"]
