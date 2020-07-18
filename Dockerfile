FROM golang:alpine

MAINTAINER Shivansh Kuchchal

RUN mkdir /hedron
WORKDIR /hedron

COPY . .
RUN go build -o ./bin/hedron-bin ./main.go

CMD ["/hedron/scripts/run_prod.sh"]