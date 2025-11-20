package cmd

import (
	"fmt"
	"os"

	"github.com/UltiRequiem/kimera/core"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "kimera",
		Short: "A JavaScript runtime powered by QuickJS",
		Long:  "Kimera is a modern JavaScript/TypeScript runtime that provides a REPL and file execution capabilities.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return core.Repl()
		},
		SilenceUsage: true,
	}

	rootCmd.AddCommand(newVersionCommand())
	rootCmd.AddCommand(newRunCommand())

	return rootCmd
}

func newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version information",
		Long:  "Display the current version of Kimera.",
		Run: func(cmd *cobra.Command, args []string) {
			core.PrintVersion()
		},
	}
}

func newRunCommand() *cobra.Command {
	opts := &core.RunOptions{}

	runCmd := &cobra.Command{
		Use:   "run [file]",
		Short: "Run a JavaScript or TypeScript file",
		Long:  "Execute a JavaScript or TypeScript file with optional permission flags for filesystem, network, and environment variable access.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.FilePath = args[0]

			if err := core.RunFile(*opts); err != nil {
				return fmt.Errorf("failed to run file: %w", err)
			}

			return nil
		},
		SilenceUsage: true,
	}

	runCmd.Flags().BoolVar(&opts.AllowFS, "fs", false, "Allow filesystem access")
	runCmd.Flags().BoolVar(&opts.AllowNet, "net", false, "Allow network access")
	runCmd.Flags().BoolVar(&opts.AllowEnv, "env", false, "Allow environment variable access")

	return runCmd
}

func Execute() error {
	rootCmd := NewRootCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	return nil
}
