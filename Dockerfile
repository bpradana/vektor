# build stage
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o vektor ./main.go

# deploy stage
FROM golang:1.19-alpine
WORKDIR /app
COPY --from=builder /app/vektor .
COPY --from=builder /app/config/config.json /app/config/config.json
RUN mkdir /app/data
CMD ["./vektor"]
