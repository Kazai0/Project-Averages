FROM golang:latest

LABEL maintainer "Jean Franzen <jean.franzen@hotmail.com>"

WORKDIR /app

COPY go.mod go.sum./

COPY . . 

RUN go build main.go

EXPOSE 3000

CMD ["./main"]