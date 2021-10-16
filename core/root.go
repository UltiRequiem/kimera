package core

import (
	"os"
        _ "embed"
	"runtime"

	"github.com/lithdew/quickjs"
)

//go:embed std/*
var codeGlobals string

func RunFile(fileToRun string) {

	code, errorReadingFile := os.ReadFile(fileToRun)

	CheckJSError(errorReadingFile)

	// Ensure that always operates in the exact same thread
	runtime.LockOSThread()

	runtimejs := quickjs.NewRuntime()
	defer runtimejs.Free()
	ctx := runtimejs.NewContext()
	defer ctx.Free()

	globals := ctx.Globals()
	globals.Set("__dispatch", ctx.Function(Globals))

	k, errorInjectingGlobals := ctx.Eval(codeGlobals)
	CheckJSError(errorInjectingGlobals)
	defer k.Free()

	result, error := ctx.EvalFile(string(code), "s")

	CheckJSError(error)

	result.Free()
}
