package core

import (
	"fmt"
	"os"

	"github.com/lithdew/quickjs"
)

// handleGetEnv handles reading environment variables
func handleGetEnv(ctx *quickjs.Context, args []quickjs.Value, permissions PermissionContext) quickjs.Value {
	if !permissions.AllowEnv {
		return ctx.ThrowError(fmt.Errorf("environment variable access denied. Use --env flag to allow"))
	}
	if len(args) < 2 {
		return ctx.ThrowTypeError("getEnv requires a variable name")
	}
	varName := args[1].String()
	value := os.Getenv(varName)
	return ctx.String(value)
}

// handleSetEnv handles setting environment variables
func handleSetEnv(ctx *quickjs.Context, args []quickjs.Value, permissions PermissionContext) quickjs.Value {
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
}
