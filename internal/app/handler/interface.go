package handler

import "net/http"

// HandleInterface interface for all handlers
type HandleInterface interface {
	Test() http.HandlerFunc
	GetTempToken() http.HandlerFunc
}
