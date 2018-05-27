package timer

import (
	"github.com/IaaS/consul"
	"github.com/jasonlvhit/gocron"
)

func Start() {
	s := gocron.NewScheduler()
	s.Every(2).Seconds().Do(consul.RegisterServer)
	go s.Start()
}
