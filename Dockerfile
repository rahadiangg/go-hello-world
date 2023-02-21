FROM golang:1.18-alpine as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o go-hello-world .

###############

# scratch bisa dibilang image kosong
# Bisa juga menggunakan image alpine:latest
FROM scratch
ENV APP_PORT=3000

WORKDIR /app
COPY --from=builder /app/go-hello-world .

ENTRYPOINT [ "./go-hello-world" ]

