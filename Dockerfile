FROM golang:alpine

RUN mkdir /hedron

WORKDIR /hedron

COPY go.mod .
COPY go.sum .

COPY . .

RUN go build -o hedron-bin ./main.go

RUN source ./cmd/configs.sh

EXPOSE $PORT

CMD ["/hedron/hedron-bin"]




