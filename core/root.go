package core

import (
	_ "embed"
	"os"
	"runtime"

	"github.com/lithdew/quickjs"
)

//go:embed std/*
var codeGlobals string

func RunFile(fileToRun string) {

	code, errorReadingFile := os.ReadFile(fileToRun)

	CheckJSError(errorReadingFile, true)

	// Ensure that always operates in the exact same thread
	runtime.LockOSThread()

	runtimejs := quickjs.NewRuntime()
	defer runtimejs.Free()
	ctx := runtimejs.NewContext()
	defer ctx.Free()

	globals := ctx.Globals()
	globals.Set("__dispatch", ctx.Function(Globals))

	k, errorInjectingGlobals := ctx.Eval(codeGlobals)
	CheckJSError(errorInjectingGlobals, true)
	defer k.Free()

	result, error := ctx.EvalFile(string(code), "s")

	CheckJSError(error, true)

	result.Free()
}
