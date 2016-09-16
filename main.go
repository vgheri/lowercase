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
	// Intercepting function which injects the request ID into the request context.
	option := httptransport.ServerBefore(setRequestIDInContext(),
		setClientIPInContext())
	var svc LowercaseService
	svc = lowercaseService{}
	svc = loggingMiddleware(logger)(svc)

	lowercase := makeLowercaseEndpoint(svc)
	lowercase = transportLoggingMiddleware(log.NewContext(logger).With("method", "lowercase"))(lowercase)

	lowercaseHandler := httptransport.NewServer(
		ctx,
		lowercase,
		decodeLowercaseRequest,
		encodeResponse,
		option,
	)

	http.Handle("/", lowercaseHandler)
	logger.Log(http.ListenAndServe(":1338", nil))
}
