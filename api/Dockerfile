FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

ENV USER=user
ENV DB_NAME=dbname
ENV PASSWORD=password

CMD ["./main"]
