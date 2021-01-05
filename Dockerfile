FROM golang

RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/mysql

WORKDIR /home/gamemanagement/application
