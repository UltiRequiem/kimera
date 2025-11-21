package core

import (
	"github.com/UltiRequiem/kimera/core/handlers"
	"github.com/UltiRequiem/kimera/core/types"
	"github.com/lithdew/quickjs"
)

func MakeGlobals(permissions types.PermissionContext) func(*quickjs.Context, quickjs.Value, []quickjs.Value) quickjs.Value {
	return func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		return Globals(ctx, this, args, permissions)
	}
}

func Globals(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value, permissions types.PermissionContext) quickjs.Value {
	switch args[0].String() {
	case "console":
		return handlers.Console(ctx, args)
	case "close":
		return handlers.Close(ctx, args)
	case "readFile":
		return handlers.ReadFile(ctx, args, permissions)
	case "writeFile":
		return handlers.WriteFile(ctx, args, permissions)
	case "fetch":
		return handlers.Fetch(ctx, args, permissions)
	case "getEnv":
		return handlers.GetEnv(ctx, args, permissions)
	case "setEnv":
		return handlers.SetEnv(ctx, args, permissions)
	case "httpCreateServer":
		return handlers.HTTPCreateServer(ctx, args, permissions)
	case "httpServerListen":
		return handlers.HTTPServerListen(ctx, args, permissions)
	case "httpServerClose":
		return handlers.HTTPServerClose(ctx, args, permissions)
	}

	return ctx.Null()
}
