FROM golang:1.17-alpine AS build
WORKDIR /golang/keng-golang/homework/3.1_docker
#LABEL multi.label1="value1" multi.label2="value2" other="value3"
COPY /golang/keng-golang/ /httpserver
EXPOSE 80
ENTRYPOINT ["/httpserver"]