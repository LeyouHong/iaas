package consul

import (
	"fmt"
	"github.com/IaaS/conf"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"strconv"
)

func RegisterServer() {
	fmt.Println("RegisterServer")
	if CheckService(conf.Conf.Consul.Cli, conf.Conf.Service.Hostname) {
		return
	}

	port, err := strconv.Atoi(conf.Conf.Service.Port)
	if err != nil {
		log.Fatal(err)
	}

	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = conf.Conf.Service.Id
	registration.Name = conf.Conf.Service.Hostname
	registration.Port = port
	registration.Tags = []string{"Tags"}
	registration.Address = conf.Conf.Service.Addr
	registration.Check = &consulapi.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%s%s", registration.Address, conf.Conf.Service.CheckPort, "/check"),
		Timeout:                        "3s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "1s",
	}

	err = conf.Conf.Consul.Cli.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("register server error : ", err)
	}

	_, err = conf.Conf.Consul.Cli.KV().Put(conf.Conf.Consul.KV, nil)
	if err != nil {
		panic(err)
	}
}

func DeRegisterServer() {
	conf.Conf.Consul.Cli.Agent().ServiceDeregister(conf.Conf.Service.Id)
	conf.Conf.Consul.Cli.KV().Delete(conf.Conf.Service.Hostname, nil)
}

func CheckService(client *consulapi.Client, hostName string) bool {
	services, err := client.Agent().Checks()
	if err != nil {
		return true
	}

	// if you already register to consul and service's status is passing. you need't to register
	for _, service := range services {
		if hostName == service.ServiceName && "passing" == service.Status {
			fmt.Println("already have this service")
			return true
		}
	}

	return false
}
