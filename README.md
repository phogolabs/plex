# plex

[![Documentation][godoc-img]][godoc-url]
![License][license-img]
[![Coverage][codecov-img]][codecov-url]
[![Go Report Card][report-img]][report-url]

A multiplexer that allows GRPC and HTTP server listening on the same port

## Installation

```console
$ go get -u github.com/phogolabs/plex
```
## Getting started

```golang
import (
  "github.com/phogolabs/plex"
)

func main() {
	// create the plex server
	server := plex.NewServer(
		plex.WithAddress(":8080"),
	)

	log.Infof("server is listening on %v for grpc or http", server.Address())
	if err := server.ListenAndServe(); err != nil {
		log.WithError(err).Error("server listen and serve failed")
	}
}
```

## Contributing

We are open for any contributions. Just fork the
[project](https://github.com/phogolabs/log).

[report-img]: https://goreportcard.com/badge/github.com/phogolabs/plex
[report-url]: https://goreportcard.com/report/github.com/phogolabs/plex
[codecov-url]: https://codecov.io/gh/phogolabs/plex
[codecov-img]: https://codecov.io/gh/phogolabs/plex/branch/master/graph/badge.svg
[action-img]: https://github.com/phogolabs/plex/workflows/main/badge.svg
[action-url]: https://github.com/phogolabs/plex/actions
[log-url]: https://github.com/phogolabs/plex
[godoc-url]: https://godoc.org/github.com/phogolabs/plex
[godoc-img]: https://godoc.org/github.com/phogolabs/plex?status.svg
[license-img]: https://img.shields.io/badge/license-MIT-blue.svg
[software-license-url]: LICENSE
