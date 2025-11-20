package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/lithdew/quickjs"
)

// handleHTTPCreateServer handles HTTP server creation
func handleHTTPCreateServer(ctx *quickjs.Context, args []quickjs.Value, permissions PermissionContext) quickjs.Value {
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
}

// handleHTTPServerListen handles starting an HTTP server
func handleHTTPServerListen(ctx *quickjs.Context, args []quickjs.Value, permissions PermissionContext) quickjs.Value {
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
}

// handleHTTPServerClose handles closing an HTTP server
func handleHTTPServerClose(ctx *quickjs.Context, args []quickjs.Value, permissions PermissionContext) quickjs.Value {
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
