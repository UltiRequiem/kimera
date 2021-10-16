package cmd

import (
	"fmt"

	"github.com/lithdew/quickjs"
)

func Globals(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {

	switch args[0].String() {

	case "console":
		fmt.Println(args[1].String())
	}

	return ctx.Null()
}
