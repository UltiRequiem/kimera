package core

import (
	"fmt"
	"os"
	"runtime"

	"github.com/lithdew/quickjs"
)

func RunFile(fileToRun string) {

	code, errorReadingFile := os.ReadFile(fileToRun)

	if errorReadingFile != nil {
		fmt.Println(errorReadingFile)
                os.Exit(1)
	}

	// Ensure that always operates in the exact same thread
	runtime.LockOSThread()

	runtimejs := quickjs.NewRuntime()
	defer runtimejs.Free()

	context := runtimejs.NewContext()
	defer context.Free()

	result, error := context.EvalFile(string(code), "s")

	if error != nil {
		fmt.Printf("Error while running %s.\n", fileToRun)
		fmt.Println(error)
	}

	result.Free()
}
