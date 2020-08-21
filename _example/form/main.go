package main

import (
	"github.com/phogolabs/log"
	"github.com/phogolabs/plex/_example/form/service"
	"github.com/phogolabs/plex/grpc"
	"github.com/phogolabs/plex/http"
	"go.opentelemetry.io/otel/exporters/stdout"
)

func main() {
	pusher, err := stdout.InstallNewPipeline(nil, nil)
	if err != nil {
		log.WithError(err).Fatal("cannot install pipeline")
	}

	defer pusher.Stop()
	// create the plex server
	server := service.New()

	http.NewHeartbeat().
		Mount(server.Proxy)

	grpc.NewHeartbeat().
		Mount(server.Gateway)

	log.Infof("server is listening on %v for grpc or http", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.WithError(err).Error("server listen and serve failed")
	}
}
