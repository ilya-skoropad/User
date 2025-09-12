package middleware

import (
	"log"
	"net/http"
	"time"
)

type LoggerMiddleware struct {
	nextHandler http.Handler
	Logger      *log.Logger
}

func (l *LoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.nextHandler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

func NewLoggerMiddleware(nextHandler http.Handler, logger *log.Logger) *LoggerMiddleware {
	return &LoggerMiddleware{
		nextHandler: nextHandler,
		Logger:      logger,
	}
}
