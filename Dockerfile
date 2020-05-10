FROM golang:alpine

RUN mkdir /hedron

WORKDIR /hedron

COPY go.mod .
COPY go.sum .
COPY . .

RUN go build -o hedron-bin ./main.go

CMD ["/hedron/hedron-bin"]




