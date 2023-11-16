// pkg/middleware/middleware.go
package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourproject/pkg/logger"
)

type responseWriter struct {
	gin.ResponseWriter
	Writer io.Writer
	Body   *bytes.Buffer
}

func (w *responseWriter) Write(b []byte) (int, error) {
	if w.Body != nil {
		w.Body.Write(b)
	}
	return w.ResponseWriter.Write(b)
}

func RequestLoggerMiddleware(log *logger.CustomLogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		var buf bytes.Buffer
		writer := &responseWriter{ResponseWriter: ctx.Writer, Writer: &buf, Body: &buf}

		ctx.Writer = writer

		ctx.Next()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()

		switch {
		case statusCode < 300:
			log.Info("method: %s, uri: %s, status: %d, latency: %s, response: %s", reqMethod, reqUri, statusCode, latencyTime, buf.String())
		case statusCode < 500:
			log.Error("method: %s, uri: %s, status: %d, latency: %s, response: %s", reqMethod, reqUri, statusCode, latencyTime, buf.String())
		default:
			log.Fatal("method: %s, uri: %s, status: %d, latency: %s, response: %s", reqMethod, reqUri, statusCode, latencyTime, buf.String())
		}
	}
}
