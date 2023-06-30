FROM ubuntu:22.04 as base

ARG EMAIL=email@example.org
# Override this email with your MyBMW email

ARG PASS=example
# Override this passwrod with your MyBMW password

RUN apt update -y

RUN apt install python3-pip golang-go jq curl -y

RUN pip3 install --upgrade bimmer_connected

RUN mkdir -p /app

WORKDIR /app

ADD . /app

ENV GO111MODULE=on

RUN go mod init download

RUN go get ./...

RUN go build ./bmw_exporter.go

CMD ["/app/bmw_exporter"]