package handlers

import (
	"fmt"
	"os"

	"github.com/lithdew/quickjs"
)

// Console handles console output
func Console(ctx *quickjs.Context, args []quickjs.Value) quickjs.Value {
	fmt.Println(args[1].String())
	return ctx.Null()
}

// Close handles application exit
func Close(ctx *quickjs.Context, args []quickjs.Value) quickjs.Value {
	os.Exit(1)
	return ctx.Null()
}
