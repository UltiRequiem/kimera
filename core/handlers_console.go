package core

import (
	"fmt"
	"os"

	"github.com/lithdew/quickjs"
)

// handleConsole handles console output
func handleConsole(ctx *quickjs.Context, args []quickjs.Value) quickjs.Value {
	fmt.Println(args[1].String())
	return ctx.Null()
}

// handleClose handles application exit
func handleClose(ctx *quickjs.Context, args []quickjs.Value) quickjs.Value {
	os.Exit(1)
	return ctx.Null()
}
