package core

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/lithdew/quickjs"
)

func Repl() error {
	fmt.Printf("Kimera %s\n", VERSION)
	fmt.Println("Exit using Ctrl+C or Ctrl+D")
	fmt.Println()

	jsRuntime := quickjs.NewRuntime()
	defer jsRuntime.Free()

	ctx := jsRuntime.NewContext()
	defer ctx.Free()

	// Inject globals
	globals := ctx.Globals()
	globals.Set("__dispatch", ctx.Function(Globals))

	globalsResult, err := ctx.Eval(codeGlobals)
	if err != nil {
		return fmt.Errorf("failed to inject globals: %w", err)
	}
	globalsResult.Free()

	reader := bufio.NewReader(os.Stdin)
	buffer := ""

	for{
		fmt.Print("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("\nGoodbye!")
				return nil
			}
			return fmt.Errorf("failed to read input: %w", err)
		}

		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		output, evalErr := evalLine(ctx, line, &buffer)
		if evalErr != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", evalErr)
			buffer = "" // Reset buffer on error
			continue
		}

		if output != "" && output != "undefined" {
			fmt.Println(output)
		}
	}
}

func evalLine(ctx *quickjs.Context, line string, buffer *string) (string, error) {
	fullCode := *buffer + line

	result, err := ctx.Eval(fullCode)
	if err != nil {
		if *buffer != "" {
			*buffer += line + "\n"
			return "", fmt.Errorf("syntax error (use empty line to reset): %w", err)
		}
		return "", err
	}
	defer result.Free()

	*buffer = ""

	return result.String(), nil
}
