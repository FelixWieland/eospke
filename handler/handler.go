package handler

import (
	"net/http"
	"sync"
)

var (
	errBadRequest = newHTTPError(http.StatusBadRequest, "bad request")
)

func newHTTPError(code int, msg string) *HTTPError {
	return &HTTPError{code: code, msg: msg}
}

type Handler struct {
	m sync.Mutex
}
