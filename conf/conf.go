package conf

import (
	"github.com/IaaS/util"
	"github.com/Sirupsen/logrus"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"os"
)

type ServiceConf struct {
	Hostname  string
	CheckPort string
	Id        string
	Addr      string
	Port      string
}

type ConsulConf struct {
	Addr string
	Cli  *consulapi.Client
	KV   *consulapi.KVPair
}

var Conf = &Config{&ServiceConf{
	os.Getenv("SERVICE_NAME"),
	os.Getenv("SERVICE_CHECKPORT"),
	os.Getenv("SERVICE_ID"),
	util.GetIntranetIp(),
	os.Getenv("SERVICE_PORT"),
}, &ConsulConf{
	os.Getenv("CONSUL_ADDRESS"),
	nil,
	nil,
}, nil,
}

type Config struct {
	Service *ServiceConf
	Consul  *ConsulConf
	Logger  *logrus.Logger
}

func (conf *Config) NewConf() {
	config := consulapi.DefaultConfig()
	config.Address = Conf.Consul.Addr
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul new client error : ", err)
	}

	Conf.Consul.Cli = client
	Conf.Logger = &logrus.Logger{}

	Conf.Consul.KV = &consulapi.KVPair{
		Key:   Conf.Service.Hostname,
		Value: []byte(Conf.Service.Addr + ":" + Conf.Service.Port),
	}
}
