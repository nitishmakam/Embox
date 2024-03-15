FROM golang:1.22-alpine3.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:3.19
RUN addgroup app && adduser -S -G app app
USER app

WORKDIR /app
COPY --from=builder --chown=app:app app/main ./

EXPOSE 8443

CMD ["./main"]