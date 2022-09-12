FROM golang:alpine

RUN mkdir /UNcademy_account_ms
RUN go install github.com/beego/bee/v2@latest

WORKDIR /UNcademy_account_ms

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

EXPOSE 8001

CMD ["bee", "run"]