FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apt-get update && \
apt-get install -y curl && \
apt-get install -y protobuf-compiler && \
apk --no-cache add ca-certificates

WORKDIR /root/

RUN mkdir migrations

COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./main"]