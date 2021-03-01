package main

import (
	"github.com/phogolabs/log"
	"github.com/phogolabs/plex/example/internal/service"
	"github.com/phogolabs/plex/service/health"
)

func main() {
	// create the plex server
	server := service.New()

	// mount the health checker
	checker := health.New()
	checker.Mount(server)

	log.Infof("server is listening on %v for grpc or http", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.WithError(err).Error("server listen and serve failed")
	}
}
