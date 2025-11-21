package handlers

import (
	"fmt"
	"os"

	"github.com/UltiRequiem/kimera/core/types"
	"github.com/lithdew/quickjs"
)

// GetEnv handles reading environment variables
func GetEnv(ctx *quickjs.Context, args []quickjs.Value, permissions types.PermissionContext) quickjs.Value {
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

// SetEnv handles setting environment variables
func SetEnv(ctx *quickjs.Context, args []quickjs.Value, permissions types.PermissionContext) quickjs.Value {
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
