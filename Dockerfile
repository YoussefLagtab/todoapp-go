FROM golang:alpine3.19 as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# TODO: use SERVER_PROTOCOL env var
RUN go build ./cmd/todoapp-http

# ------------------------
FROM alpine:3.19

WORKDIR /app
COPY --from=build /app/todoapp-http .

CMD ["./todoapp-http"]
