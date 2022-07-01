FROM golang:1.11 AS builder

RUN mkdir /app
WORKDIR /app
COPY . .

RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build main.go

FROM alpine
RUN mkdir /app
COPY --from=builder /app/main /app/maim
