package main

import (
	"net/http"
	"os"

	log "github.com/go-kit/kit/log"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()
	logger := log.NewLogfmtLogger(os.Stderr)
	var svc LowercaseService
	svc = lowercaseService{}
	svc = loggingMiddleware(logger)(svc)

	lowercase := makeLowercaseEndpoint(svc)

	lowercaseHandler := httptransport.NewServer(
		ctx,
		lowercase,
		decodeLowercaseRequest,
		encodeResponse,
	)

	http.Handle("/", lowercaseHandler)
	logger.Log(http.ListenAndServe(":1338", nil))
}
