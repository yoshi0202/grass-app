FROM golang:1.15.2-buster
RUN go get -u -t -v github.com/gin-gonic/gin
RUN go get -u -t -v gorm.io/gorm
RUN go get -u -t -v gorm.io/driver/mysql
RUN go get -u -t -v gorm.io/datatypes
RUN go get -u -t -v github.com/joho/godotenv 
RUN go get -u -t -v github.com/gin-contrib/sessions
RUN go get -u -t -v github.com/gin-contrib/sessions/cookie
RUN go get -u -t -v github.com/gin-contrib/cors
RUN mkdir -p /go/src/github.com/yoshi0202/grass-app/app