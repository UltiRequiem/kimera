package core

import (
	"bufio"
	"fmt"
	"os"

	"github.com/lithdew/quickjs"
)

func Repl() {
	fmt.Printf("Kimera %s", VERSION)
	fmt.Println("exit using ctrl+c or close()")

	for true {
		fmt.Print("> ")

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		stringToEval := ""
		fmt.Println(Eval(text, &stringToEval))
		stringToEval += ";undefined;"
	}
}

func Eval(text string, buffer *string) string {
	runtime := quickjs.NewRuntime()
	defer runtime.Free()

	ctx := runtime.NewContext()
	defer ctx.Free()

	globalsEval := ctx.Globals()
	globalsEval.Set("__dispatch", ctx.Function(Globals))

	k, errorInjectingGlobals := ctx.Eval(codeGlobals)
	CheckJSError(errorInjectingGlobals, false)
	defer k.Free()

	result, err := ctx.Eval(*buffer + text)

	if err != nil {
		CheckJSError(err, false)
		return result.String()
	}

	*buffer += fmt.Sprintf(";undefined; %s", text)

	result.Free()

	return result.String()
}
