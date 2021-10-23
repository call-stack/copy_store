FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . /app

RUN go build  /app/cmd/read/read.go

CMD [ "./read" ]