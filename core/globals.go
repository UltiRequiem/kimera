package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/lithdew/quickjs"
)

type PermissionContext struct {
	AllowFS  bool
	AllowNet bool
	AllowEnv bool
}

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

func MakeGlobals(permissions PermissionContext) func(*quickjs.Context, quickjs.Value, []quickjs.Value) quickjs.Value {
	return func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		return Globals(ctx, this, args, permissions)
	}
}

func Globals(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value, permissions PermissionContext) quickjs.Value {

	switch args[0].String() {

	case "console":
		fmt.Println(args[1].String())
	case "close":
		os.Exit(1)
	case "readFile":
		if !permissions.AllowFS {
			return ctx.ThrowError(fmt.Errorf("filesystem access denied. Use --fs flag to allow"))
		}
		if len(args) < 2 {
			return ctx.ThrowTypeError("readFile requires a file path")
		}
		filePath := args[1].String()
		content, err := os.ReadFile(filePath)
		if err != nil {
			return ctx.ThrowError(err)
		}
		return ctx.String(string(content))
	case "writeFile":
		if !permissions.AllowFS {
			return ctx.ThrowError(fmt.Errorf("filesystem access denied. Use --fs flag to allow"))
		}
		if len(args) < 3 {
			return ctx.ThrowTypeError("writeFile requires a file path and content")
		}
		filePath := args[1].String()
		content := args[2].String()
		err := os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			return ctx.ThrowError(err)
		}
		return ctx.Null()
	case "fetch":
		if !permissions.AllowNet {
			return ctx.ThrowError(fmt.Errorf("network access denied. Use --net flag to allow"))
		}
		if len(args) < 2 {
			return ctx.ThrowTypeError("fetch requires a URL")
		}
		url := args[1].String()
		
		// Default options
		method := "GET"
		var body io.Reader
		headers := make(map[string]string)
		
		// Parse options if provided
		if len(args) >= 3 && !args[2].IsNull() && !args[2].IsUndefined() {
			optionsJSON := args[2].String()
			var options map[string]interface{}
			if err := json.Unmarshal([]byte(optionsJSON), &options); err == nil {
				if m, ok := options["method"].(string); ok {
					method = strings.ToUpper(m)
				}
				if b, ok := options["body"].(string); ok && b != "" {
					body = bytes.NewBufferString(b)
				}
				if h, ok := options["headers"].(map[string]interface{}); ok {
					for key, val := range h {
						if strVal, ok := val.(string); ok {
							headers[key] = strVal
						}
					}
				}
			}
		}
		
		// Create request
		req, err := http.NewRequest(method, url, body)
		if err != nil {
			return ctx.ThrowError(err)
		}
		
		// Set headers
		for key, val := range headers {
			req.Header.Set(key, val)
		}
		
		// Make request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return ctx.ThrowError(err)
		}
		defer resp.Body.Close()
		
		// Read response body
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return ctx.ThrowError(err)
		}
		
		// Create response object
		responseHeaders := make(map[string]string)
		for key, values := range resp.Header {
			if len(values) > 0 {
				responseHeaders[key] = values[0]
			}
		}
		
		responseObj := map[string]interface{}{
			"ok":         resp.StatusCode >= 200 && resp.StatusCode < 300,
			"status":     resp.StatusCode,
			"statusText": resp.Status,
			"headers":    responseHeaders,
			"body":       string(responseBody),
			"url":        url,
		}
		
		responseJSON, err := json.Marshal(responseObj)
		if err != nil {
			return ctx.ThrowError(err)
		}
		
		return ctx.String(string(responseJSON))
	case "getEnv":
		if !permissions.AllowEnv {
			return ctx.ThrowError(fmt.Errorf("environment variable access denied. Use --env flag to allow"))
		}
		if len(args) < 2 {
			return ctx.ThrowTypeError("getEnv requires a variable name")
		}
		varName := args[1].String()
		value := os.Getenv(varName)
		return ctx.String(value)
	case "setEnv":
		if !permissions.AllowEnv {
			return ctx.ThrowError(fmt.Errorf("environment variable access denied. Use --env flag to allow"))
		}
		if len(args) < 3 {
			return ctx.ThrowTypeError("setEnv requires a variable name and value")
		}
		varName := args[1].String()
		value := args[2].String()
		err := os.Setenv(varName, value)
		if err != nil {
			return ctx.ThrowError(err)
		}
		return ctx.Null()
	case "httpCreateServer":
		if !permissions.AllowNet {
			return ctx.ThrowError(fmt.Errorf("network access denied. Use --net flag to allow"))
		}
		if len(args) < 2 {
			return ctx.ThrowTypeError("httpCreateServer requires a handler ID")
		}
		
		// Create a channel for this server
		serverState.Lock()
		serverID := fmt.Sprintf("server_%d", serverState.nextID)
		serverState.nextID++
		serverState.channels[serverID] = make(chan ServerRequest)
		serverState.Unlock()
		
		return ctx.String(serverID)
	case "httpServerListen":
		if !permissions.AllowNet {
			return ctx.ThrowError(fmt.Errorf("network access denied. Use --net flag to allow"))
		}
		if len(args) < 4 {
			return ctx.ThrowTypeError("httpServerListen requires serverID, port, and handlerID")
		}
		
		serverID := args[1].String()
		port := args[2].String()
		handlerID := args[3].String()
		
		serverState.Lock()
		reqChan, exists := serverState.channels[serverID]
		if !exists {
			serverState.Unlock()
			return ctx.ThrowError(fmt.Errorf("server not found: %s", serverID))
		}
		serverState.Unlock()
		
		// Create HTTP handler that sends requests through the channel
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			// Read request body
			body, _ := io.ReadAll(r.Body)
			defer r.Body.Close()
			
			// Build request object
			headers := make(map[string]string)
			for key, values := range r.Header {
				if len(values) > 0 {
					headers[key] = values[0]
				}
			}
			
			req := HTTPRequest{
				Method:  r.Method,
				URL:     r.URL.String(),
				Path:    r.URL.Path,
				Query:   r.URL.RawQuery,
				Headers: headers,
				Body:    string(body),
			}
			
			// Send request and wait for response
			respChan := make(chan HTTPResponse, 1)
			reqChan <- ServerRequest{Request: req, Response: respChan}
			
			response := <-respChan
			
			// Set status code
			w.WriteHeader(response.Status)
			
			// Set headers
			for key, val := range response.Headers {
				w.Header().Set(key, val)
			}
			
			// Write body
			fmt.Fprint(w, response.Body)
		})
		
		// Create and store server
		server := &http.Server{
			Addr:    ":" + port,
			Handler: mux,
		}
		
		serverState.Lock()
		serverState.servers[serverID] = server
		serverState.Unlock()
		
		// Start server in a goroutine
		go func() {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
			}
		}()
		
		// Process requests in the main thread (where QuickJS is safe)
		for {
			select {
			case serverReq := <-reqChan:
				// Marshal request to JSON
				requestJSON, _ := json.Marshal(serverReq.Request)
				
				// Evaluate JavaScript code to call the handler
				code := fmt.Sprintf(`
					(function() {
						const request = %s;
						const response = globalThis.__serverHandlers['%s'](request);
						return JSON.stringify(response);
					})()
				`, string(requestJSON), handlerID)
				
				result, err := ctx.Eval(code)
				if err != nil {
					serverReq.Response <- HTTPResponse{
						Status: http.StatusInternalServerError,
						Headers: map[string]string{"Content-Type": "text/plain"},
						Body: fmt.Sprintf("Internal Server Error: %v", err),
					}
					continue
				}
				
				if result.IsException() {
					result.Free()
					serverReq.Response <- HTTPResponse{
						Status: http.StatusInternalServerError,
						Headers: map[string]string{"Content-Type": "text/plain"},
						Body: "Handler exception",
					}
					continue
				}
				
				// Parse response
				responseJSON := result.String()
				result.Free()
				
				var response HTTPResponse
				if err := json.Unmarshal([]byte(responseJSON), &response); err != nil {
					serverReq.Response <- HTTPResponse{
						Status: http.StatusInternalServerError,
						Headers: map[string]string{"Content-Type": "text/plain"},
						Body: "Invalid response format",
					}
					continue
				}
				
				// Default status if not set
				if response.Status == 0 {
					response.Status = http.StatusOK
				}
				
				// Send response back
				serverReq.Response <- response
			}
		}
	case "httpServerClose":
		if !permissions.AllowNet {
			return ctx.ThrowError(fmt.Errorf("network access denied. Use --net flag to allow"))
		}
		if len(args) < 2 {
			return ctx.ThrowTypeError("httpServerClose requires serverID")
		}
		
		serverID := args[1].String()
		
		serverState.Lock()
		server, exists := serverState.servers[serverID]
		if !exists {
			serverState.Unlock()
			return ctx.ThrowError(fmt.Errorf("server not found: %s", serverID))
		}
		
		// Close the channel
		if ch, ok := serverState.channels[serverID]; ok {
			close(ch)
			delete(serverState.channels, serverID)
		}
		
		delete(serverState.servers, serverID)
		serverState.Unlock()
		
		if err := server.Close(); err != nil {
			return ctx.ThrowError(fmt.Errorf("failed to close server: %w", err))
		}
		
		return ctx.Null()
	}

	return ctx.Null()
}
