package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/DavidMovas/me-ardan-labs-service/foundation/logger"
	"github.com/DavidMovas/me-ardan-labs-service/foundation/web"
	"github.com/ardanlabs/service/app/sdk/errs"
	"net/http"
	"time"
)

func Logger(logger *logger.Logger) web.MidFunc {
	return func(next web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, r *http.Request) web.Encoder {
			now := time.Now()

			path := r.URL.Path
			if r.URL.RawQuery != "" {
				path = fmt.Sprintf("%s?%s", path, r.URL.RawQuery)
			}

			logger.Info(ctx, "request started", "method", r.Method, "path", path, "remoteaddr", r.RemoteAddr)

			resp := next(ctx, r)
			err := isError(resp)

			var statusCode = errs.OK
			if err != nil {
				statusCode = errs.Internal

				var v *errs.Error
				if errors.As(err, &v) {
					statusCode = v.Code
				}
			}

			logger.Info(ctx, "request completed", "method", r.Method, "path", path, "remoteaddr", r.RemoteAddr,
				"statuscode", statusCode, "since", time.Since(now).String())

			return resp
		}

		return h
	}
}
