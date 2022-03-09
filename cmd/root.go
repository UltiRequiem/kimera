package cmd

import (
	"github.com/UltiRequiem/kimera/core"
	"github.com/spf13/cobra"
)

func main() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "Runs the REPL.",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			core.Repl()
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version.",
		Run: func(cmd *cobra.Command, args []string) {
			core.PrintVersion()
		},
	}

	var fsFlag bool
	var netFlag bool
	var envFlag bool

	var runCmd = &cobra.Command{
		Use:   "run [file]",
		Short: "Run a JavaScript file.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			core.RunFile(args[0])
		},
	}
	runCmd.Flags().BoolVar(&fsFlag, "fs", false, "Allow file system access")
	runCmd.Flags().BoolVar(&netFlag, "net", false, "Allow net access")
	runCmd.Flags().BoolVar(&envFlag, "env", false, "Allow Environment Variables access")

	rootCmd.AddCommand(runCmd, versionCmd)

	return rootCmd
}
