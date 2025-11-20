package core

import (
	_ "embed"
	"fmt"
	"os"
	"runtime"

	esbuild "github.com/evanw/esbuild/pkg/api"
	"github.com/lithdew/quickjs"
)

//go:embed std/root.js
var codeGlobals string

type RunOptions struct {
	FilePath string
	AllowFS  bool
	AllowNet bool
	AllowEnv bool
}

func RunFile(opts RunOptions) error {
	if opts.FilePath == "" {
		return fmt.Errorf("file path cannot be empty")
	}

	code, err := os.ReadFile(opts.FilePath)

	if err != nil {
		return fmt.Errorf("failed to read file %q: %w", opts.FilePath, err)
	}

	parsedCode := esbuild.Transform(string(code), esbuild.TransformOptions{
		Loader: esbuild.LoaderTS,
	})

	if len(parsedCode.Errors) > 0 {
		return fmt.Errorf("transpilation errors:\n%s", formatEsbuildErrors(parsedCode.Errors))
	}

	runtime.LockOSThread()

	jsRuntime := quickjs.NewRuntime()
	defer jsRuntime.Free()

	ctx := jsRuntime.NewContext()
	defer ctx.Free()

	globals := ctx.Globals()
	globals.Set("__dispatch", ctx.Function(Globals))

	// TODO: Use opts.AllowFS, opts.AllowNet, opts.AllowEnv to control permissions

	// Inject global code
	globalsResult, err := ctx.Eval(codeGlobals)
	
	if err != nil {
		return fmt.Errorf("failed to inject globals: %w", err)
	}
	
	globalsResult.Free()

	result, err := ctx.EvalFile(string(parsedCode.Code), opts.FilePath)
	
	if err != nil {
		return fmt.Errorf("runtime error: %w", err)
	}
	
	defer result.Free()

	return nil
}

func formatEsbuildErrors(errors []esbuild.Message) string {
	result := ""
	
	for i, err := range errors {
		if i > 0 {
			result += "\n"
		}
	
		result += fmt.Sprintf("  - %s", err.Text)
	
		if err.Location != nil {
			result += fmt.Sprintf(" at %s:%d:%d",
				err.Location.File,
				err.Location.Line,
				err.Location.Column,
			)
		}
	}
	
	return result
}
