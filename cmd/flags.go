package cmd

import "flag"

func flagsArgs() (bool, []string) {
	flag.Parse()

	return flag.NArg() > 0,flag.Args()
}
