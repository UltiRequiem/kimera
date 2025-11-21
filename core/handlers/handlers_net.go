package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/UltiRequiem/kimera/core/types"
	"github.com/lithdew/quickjs"
)

// Fetch handles HTTP fetch operations
func Fetch(ctx *quickjs.Context, args []quickjs.Value, permissions types.PermissionContext) quickjs.Value {
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
}
