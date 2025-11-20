package core

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
	servers  map[string]*http.Server
	channels map[string]chan ServerRequest
	nextID   int
}

var serverState = &ServerState{
	servers:  make(map[string]*http.Server),
	channels: make(map[string]chan ServerRequest),
}
