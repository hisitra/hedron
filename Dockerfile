FROM golang:alpine

RUN mkdir /hedron

WORKDIR /hedron

COPY go.mod .
COPY go.sum .

COPY . .

RUN go build -o hedron-bin ./main.go

RUN source ./scripts/configs.sh

EXPOSE $PORT
EXPOSE $REST_PORT

CMD ["/hedron/hedron-bin"]




