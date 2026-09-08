FROM golang:1.26.5

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /main main.go

CMD ["/main"]