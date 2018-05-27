package main

import (
	"fmt"
	"github.com/IaaS/conf"
	"github.com/IaaS/consul"
	"github.com/IaaS/timer"
	"log"
	"net/http"
)

func main() {
	conf.Conf.NewConf()             //Init Config
	timer.Start()                   //Consul register
	defer consul.DeRegisterServer() //Deregister

	//This is service logic
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, from %s", conf.Conf.Service.Hostname)
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", conf.Conf.Service.Port), nil))
}
