# start a golang base image, version 1.8
FROM golang:1.8 as builder

#disable crosscompiling 
ENV CGO_ENABLED=0

#compile linux only
ENV GOOS=linux

ADD . /go/src/github.com/IaaS
RUN cd /go/src/github.com/IaaS && make

COPY --from=builder /go/src/github.com/IaaS/build/demo /usr/local/bin/
#Service config
ENV SERVICE_NAME=DEMO
ENV SERVICE_CHECKPORT=8080
ENV SERVICE_ID=DEMO
ENV SERVICE_PORT=8080

#Consul config
ENV CONSUL_ADDRESS=172.17.0.1:8500

EXPOSE 8080
EXPOSE 443

CMD ["demo"]
