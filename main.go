package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/UltiRequiem/kimera/cmd"
)

func main() {
	program := cmd.Execute()

	if err := program.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
