package core

import (
	"fmt"
	"os"

	"github.com/lithdew/quickjs"
)

func Globals(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {

	switch args[0].String() {

	case "console":
		fmt.Println(args[1].String())
	case "close":
		os.Exit(1)
	}

	return ctx.Null()
}
