package main

import (
	"time"

	"golang.org/x/net/context"

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

func (mw appLoggingMiddleware) Lowercase(ctx context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "lowercase",
			"requestID", ctx.Value("requestID"),
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Lowercase(ctx, s)
	return
}
