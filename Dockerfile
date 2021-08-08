FROM golang:latest

RUN mkdir /GO

WORKDIR /GO

RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm
RUN go get github.com/go-xorm/xorm
RUN go get go.uber.org/zap
RUN go get github.com/stretchr/testify/assert

COPY go.mod .
COPY go.sum .