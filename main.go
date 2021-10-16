package main

import (
	_ "embed"

	"github.com/UltiRequiem/kimera/cmd"
)

//go:embed std/*
var s string

func main() {
	cmd.Main(s)
}
