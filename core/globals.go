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
	}

	return ctx.Null()
}
