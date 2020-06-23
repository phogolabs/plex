package main

import (
	"github.com/phogolabs/log"
	"github.com/phogolabs/plex/_example/form/service"
)

func main() {
	// create the plex server
	server := service.New()

	log.Infof("server is listening on %v for grpc or http", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.WithError(err).Error("server listen and serve failed")
	}
}
