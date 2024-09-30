FROM golang:1.22.5-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

ENV CGO_ENABLED=0

RUN go build -ldflags="-s -w" -o main .

FROM scratch

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]