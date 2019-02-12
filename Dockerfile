FROM golang:1.8 as builder

RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/streadway/amqp
RUN go get github.com/alecthomas/template
RUN go get github.com/swaggo/swag
RUN go get github.com/swaggo/gin-swagger
RUN go get github.com/swaggo/gin-swagger/swaggerFiles
RUN go get github.com/satori/go.uuid
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go get github.com/go-sql-driver/mysql

COPY . /go/src/


WORKDIR /go/src/flow
RUN go build -o main main.go
CMD ["./main"]
# Application image.
# FROM golang:1.8

# COPY --from=builder /go/src/trainTickets/app /usr/local/bin/app

# CMD ["/usr/local/bin/app"]
