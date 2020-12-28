#開発環境用ホートリロード
FROM golang:1.15.4-alpine as dev
WORKDIR /tmp
ENV GOBIN=/usr/local/go/bin
ENV GO111MODULE=on
ENV CGO_ENABLED=0
RUN set -eux && \
apk add --update --no-cache ca-certificates git byobu

RUN go get github.com/pilu/fresh 
RUN go get github.com/jinzhu/gorm
RUN go get -u github.com/rubenv/sql-migrate/...
WORKDIR /usr/local/go/src/one-accounts
COPY . /usr/local/go/src/one-accounts
EXPOSE 8080
#CMD ["fresh"]