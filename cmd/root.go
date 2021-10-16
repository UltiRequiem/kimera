package cmd

import (
	"github.com/UltiRequiem/kimera/core"
	"github.com/spf13/cobra"
)

func Execute() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "Runs the REPL.",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			core.Repl()
		},
	}

	var versionCmd = &cobra.Command{
		Use: "version",
                Short: "Print the version.",
		Run: func(cmd *cobra.Command, args []string) {
			core.PrintVersion()
		},
	}

	var runCmd = &cobra.Command{
		Use:   "run [file]",
		Short: "Run a JavaScript file.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			core.RunFile(args[0])
		},
	}

	rootCmd.AddCommand(runCmd, versionCmd)

	return rootCmd
}
