# BUILD STAGE
FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/bin/main .


# RUN STAGE
FROM --platform=linux/x86_64 alpine:latest

COPY --from=builder /app/bin/main .
COPY --from=builder /app/debug.env .

HEALTHCHECK  --interval=60s --timeout=10s --start-period=20s CMD curl --fail 127.0.0.1:8080/ping || exit 1

EXPOSE 8080
CMD ["/main"]
