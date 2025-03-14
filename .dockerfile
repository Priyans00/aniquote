FROM golang:1.21

WORKDIR /app

COPY . .

Run go mod tidy

RUN go build -o server 

EXPOSE 3000

CMD ["./main"]