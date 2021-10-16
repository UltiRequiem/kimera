package cmd

import (
	"os"
	"runtime"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/lithdew/quickjs"
)

func Main(codeGlobals string) {
	runtime.LockOSThread()

	fileToRun := os.Args[1:][0]

	jsruntime := quickjs.NewRuntime()
	defer jsruntime.Free()

	context := jsruntime.NewContext()
	defer context.Free()

	globals := context.Globals()
	globals.Set("__dispatch", context.Function(Globals))

	k, errorInjectingGlobals := context.Eval(codeGlobals)
	CheckJSError(errorInjectingGlobals)
	defer k.Free()

	bundle := api.Build(api.BuildOptions{
		EntryPoints: []string{fileToRun},
		Outfile:     "output.js",
		Bundle:      true,
		Target:      api.ES2015,
		Write:       true,
		LogLevel:    api.LogLevelInfo,
	})
	defer os.Remove("output.js")

	result, errorEvaluatingFile := context.EvalFile(string(bundle.OutputFiles[0].Contents[:]), "s")
	CheckJSError(errorEvaluatingFile)
	result.Free()
}
