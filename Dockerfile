FROM golang:1.23.0
WORKDIR /app

COPY go.mod ./

RUN go mod download  && go mod verify

COPY . .

RUN go build  -v -o main .

CMD ["/app/main"]