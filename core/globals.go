package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/lithdew/quickjs"
)

func Globals(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {

	switch args[0].String() {

	case "console":
		fmt.Println(args[1].String())
	case "close":
		os.Exit(1)
	case "readFile":
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

	return ctx.Null()
}
