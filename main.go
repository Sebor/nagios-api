package main

import (
	"github.com/Sebor/nagios-api/api"
	"github.com/Sebor/nagios-api/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	conf := config.GetConfig()
	api := api.NewAPI(conf.Addr, conf.ObjectCacheFile, conf.CommandFile, conf.StatusFile)

	err := api.Run()
	if err != nil {
		log.Fatal(err)
	}
}
