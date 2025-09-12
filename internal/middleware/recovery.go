package middleware

import (
	"fmt"
	"net/http"
)

type recovery struct {
	nextHandler http.Handler
}

func (h *recovery) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()

		if err != nil {
			w.Header().Set("Connection", "close")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Errorf("%s", err).Error()))
		}
	}()

	h.nextHandler.ServeHTTP(w, r)
}

func NewRecovery(nextHandler http.Handler) http.Handler {
	return &recovery{
		nextHandler: nextHandler,
	}
}
