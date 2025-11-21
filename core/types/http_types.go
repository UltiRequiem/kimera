package types

import (
	"net/http"
	"sync"
)

// HTTPRequest represents an HTTP request
type HTTPRequest struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Path    string            `json:"path"`
	Query   string            `json:"query"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

// HTTPResponse represents an HTTP response
type HTTPResponse struct {
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

// ServerRequest wraps a request with a response channel
type ServerRequest struct {
	Request  HTTPRequest
	Response chan HTTPResponse
}

// ServerState holds the state of HTTP servers
type ServerState struct {
	sync.Mutex
	Servers  map[string]*http.Server
	Channels map[string]chan ServerRequest
	NextID   int
}

// GetServerState returns the global server state instance
func GetServerState() *ServerState {
	return serverState
}

var serverState = &ServerState{
	Servers:  make(map[string]*http.Server),
	Channels: make(map[string]chan ServerRequest),
}
