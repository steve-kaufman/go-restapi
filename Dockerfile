FROM golang:1.14

WORKDIR /go/src/app
COPY . .

ENV DB_PASSWORD=${DB_PASSWORD}

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]