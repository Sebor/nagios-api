package main

import (
	"github.com/sebor/nagios-api/api"
	"github.com/sebor/nagios-api/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	conf := config.GetConfig()
	api := api.NewApi(conf.Addr, conf.ObjectCacheFile, conf.CommandFile, conf.StatusFile)

	err := api.Run()
	if err != nil {
		log.Fatal(err)
	}
}
