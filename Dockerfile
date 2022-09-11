FROM golang:alpine

RUN mkdir /UNcademy_account_ms

WORKDIR /UNcademy_account_ms

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8000

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main