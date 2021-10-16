package cmd

import (
	"errors"
	"fmt"

	"github.com/lithdew/quickjs"
)

func CheckJSError(err error) {
	if err != nil {
		var evalErr *quickjs.Error

		if errors.As(err, &evalErr) {
			fmt.Println(evalErr.Cause)
			fmt.Println(evalErr.Stack)
		}

		panic(err)
	}
}
