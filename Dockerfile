FROM golang:1.15.5-alpine3.12 as build

COPY . /go/src/github.com/p4inty/podo-wundsam
WORKDIR /go/src/github.com/p4inty/podo-wundsam

RUN apk add git nodejs npm && \
    go get -d -v ./... && \
    go install -v ./... && \
    go build

WORKDIR /go/src/github.com/p4inty/podo-wundsam/assets

RUN npm install && \
    npm run production && \
    rm -rf ../assets
    rm -rf ../.gitignore


FROM alpine:latest
WORKDIR /root/
COPY --from=build /go/src/github.com/p4inty/podo-wundsam .
ENV GIN_MODE=release
CMD ["./podo-wundsam"]
EXPOSE 8080