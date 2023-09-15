# build stage
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./main.go

# deploy stage
FROM golang:1.19-alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/config/*.json .
RUN mkdir /app/temp
CMD ["./main"]
