FROM golang:1.20 as builder

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    PORT=8811

COPY *.go ./
COPY go.mod go.sum ./
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o bookstore-domain-warehouse *.go


# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY conf/*.toml conf/
COPY --from=builder /app/bookstore-domain-warehouse .

EXPOSE $PORT

# Run
CMD [ "./bookstore-domain-warehouse" ]

