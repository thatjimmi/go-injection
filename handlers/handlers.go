package handlers

import (
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

// HelloHandler is a handler that says hello.
type HelloHandler struct {
	log *zap.Logger
}

func (*HelloHandler) Pattern() string {
	return "/hello"
}

func NewHelloHandler(log *zap.Logger) *HelloHandler {
	return &HelloHandler{log: log}
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error("Error reading body", zap.Error(err))
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}

	if _, err := fmt.Fprintf(w, "Hello, %s!", body); err != nil {
		h.log.Error("Error writing response", zap.Error(err))
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}

	h.log.Info("Said hello", zap.String("body", string(body)))
}

// EchoHandler is a handler that echoes the request body.
type EchoHandler struct {
	log *zap.Logger
}

func (*EchoHandler) Pattern() string {
	return "/echo"
}

func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{log: log}
}

func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		h.log.Warn("Error copying body", zap.Error(err))
	}
}
