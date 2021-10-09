FROM golang:1.17.1

RUN mkdir /app

WORKDIR /go/src/app

RUN go mod init && \
  go get github.com/go-sql-driver/mysql && \
  go get -v github.com/rubenv/sql-migrate/... && \
  go get github.com/dgrijalva/jwt-go

CMD ["go", "run", "main.go"]
# CMD ["sql-migrate", "up"]