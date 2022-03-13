FROM golang:alpine AS devlopement
MAINTAINER akazwz
WORKDIR /home/imagin
ADD . /home/imagin
ENV GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
RUN go build -o app

FROM alpine:latest AS production
WORKDIR /root/
COPY --from=devlopement /home/imagin/app .
EXPOSE 9000:9000
ENV GIN_MODE=release
ENTRYPOINT ["./app"]