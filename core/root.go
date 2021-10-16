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
	context := runtimejs.NewContext()
	defer context.Free()

	globals := context.Globals()
	globals.Set("__dispatch", context.Function(Globals))

	k, errorInjectingGlobals := context.Eval(codeGlobals)
	CheckJSError(errorInjectingGlobals)
	defer k.Free()

	result, error := context.EvalFile(string(code), "s")

	CheckJSError(error)

	result.Free()
}
