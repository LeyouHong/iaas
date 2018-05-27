# Install
```
$ git clone ...
$ go get github.com/hashicorp/consul
$ make 
& make docker
```
# Run in docker
```
docker run -it demo
```

# Leaner Docker Image

Here's a two stage example of how we can containerize the code that will run on our nodes

First we compile our Go code in a full golang1.8 container.  Then, we copy
the executable, make sure it's runnable and buld the the much smaller image for 
deployment.   

Note this example shows how we'll make https calls from within a container.


<b>How to:</b>

First Build the executable
```
docker build -f Dockerfile.build -t iaas-build:latest .
```
Next get the exeutable from the built image
```
docker run iaas-build cat /go/src/iaassample/iaassample > iaassample
```
Now make sure it's executable
```
chmod +x iaassample
```
Create the much smaller (and faster) production image
```
docker build -f Dockerfile.production -t iaassample:latest .
```
Run it
```
docker run -it iaassample
 ```

# consul agent in docker

Get consul image
```
docker pull consul
```
Run the consul
```
docker run -d --name=dev-consul -e CONSUL_BIND_INTERFACE=docker0 --net=host consul
```
check the log and get the ip address.
```
docker logs -f dev-consul
```
type ip:port into the browser. we can access the website : http://172.17.0.1:8500/ui/#/dc1/services

run the demo by using run_in_docker.sh, then register the service.

```
curl -X PUT -d '{"id": "demo","name": "demo1","address": "172.17.0.2","port": 8080,"tags": ["172.17.0.2:8080"], "Check":{"DeregisterCriticalServiceAfter": "90m","HTTP": "http://172.17.0.2:8080/123","Interval": "10s"}}' http://172.17.0.1:8500/v1/agent/service/register
```
