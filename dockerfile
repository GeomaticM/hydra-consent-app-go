FROM golang:1.9.2-alpine

WORKDIR /go/src/hydra-consent-app-go
COPY ./ .
COPY ven vendor

#设置环境变量，使go在编译过程中从vendor目录查找依赖
ENV GO15VENDOREXPERIMENT 1

RUN go-wrapper install

EXPOSE 3000
CMD ["hydra-consent-app-go"]
