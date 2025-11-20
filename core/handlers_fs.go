package core

import (
	"fmt"
	"os"

	"github.com/lithdew/quickjs"
)

// handleReadFile handles file reading operations
func handleReadFile(ctx *quickjs.Context, args []quickjs.Value, permissions PermissionContext) quickjs.Value {
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
}

// handleWriteFile handles file writing operations
func handleWriteFile(ctx *quickjs.Context, args []quickjs.Value, permissions PermissionContext) quickjs.Value {
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
}
