package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "packrat",
	Short:         "pacman wrapper written in Go",
	Long:          "Wrapper that offers human-readable ways to run pacman package manager commands on an Arch-based system. Includes installing, updating, pruning, etc.",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error", err)
		os.Exit(1)
	}
}
