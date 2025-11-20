package core

import (
	"github.com/lithdew/quickjs"
)

func MakeGlobals(permissions PermissionContext) func(*quickjs.Context, quickjs.Value, []quickjs.Value) quickjs.Value {
	return func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		return Globals(ctx, this, args, permissions)
	}
}

func Globals(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value, permissions PermissionContext) quickjs.Value {
	switch args[0].String() {
	case "console":
		return handleConsole(ctx, args)
	case "close":
		return handleClose(ctx, args)
	case "readFile":
		return handleReadFile(ctx, args, permissions)
	case "writeFile":
		return handleWriteFile(ctx, args, permissions)
	case "fetch":
		return handleFetch(ctx, args, permissions)
	case "getEnv":
		return handleGetEnv(ctx, args, permissions)
	case "setEnv":
		return handleSetEnv(ctx, args, permissions)
	case "httpCreateServer":
		return handleHTTPCreateServer(ctx, args, permissions)
	case "httpServerListen":
		return handleHTTPServerListen(ctx, args, permissions)
	case "httpServerClose":
		return handleHTTPServerClose(ctx, args, permissions)
	}

	return ctx.Null()
}
