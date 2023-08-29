FROM golang:1.18.1

RUN go version
ENV GOPATH=/

COPY . .

RUN apt-get update
RUN apt-get -y install postgresql-client
RUN chmod +x wait-for-postgres.sh


RUN go mod download
RUN go build -o chat-app ./cmd/main.go

CMD ["./chat-app"]