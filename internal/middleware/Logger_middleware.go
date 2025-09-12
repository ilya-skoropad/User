package middleware

import (
	"log"
	"net/http"
	"time"
)

type LoggerMiddleware struct {
	nextHandler http.HandlerFunc
	Logger      *log.Logger
}

func (l *LoggerMiddleware) handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.nextHandler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

func NewLoggerMiddleware(nextHandler http.HandlerFunc, logger *log.Logger) *LoggerMiddleware {
	return &LoggerMiddleware{
		nextHandler: nextHandler,
		Logger:      logger,
	}
}
