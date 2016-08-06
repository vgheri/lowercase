package main

import (
	"time"

	log "github.com/go-kit/kit/log"
)

type appLoggingMiddleware struct {
	logger log.Logger
	next   LowercaseService
}

func loggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next LowercaseService) LowercaseService {
		return appLoggingMiddleware{logger, next}
	}
}

func (mw appLoggingMiddleware) Lowercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "lowercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Lowercase(s)
	return
}
