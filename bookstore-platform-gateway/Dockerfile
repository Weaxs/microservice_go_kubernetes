FROM golang:1.20 as builder

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    PORT=8810

COPY *.go ./
COPY go.mod go.sum ./
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o bookstore-platform-gateway *.go


# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/bookstore-platform-gateway .

EXPOSE $PORT

# Run
CMD [ "./bookstore-platform-gateway" ]

