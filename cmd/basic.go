package cmd

import (
	run "github.com/coinedd1/packrat/internal"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install <package(s)>",
	Aliases: []string{"i", "in", "inst", "-S"},
	Short:   "Install a package from the Arch Linux repository",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		pacArgs := append([]string{"-S"}, args...)
		return run.SudoStdTerminal(pacArgs...)
	},
}

var updateCmd = &cobra.Command{
	Use:     "update <package(s)>",
	Aliases: []string{"upd", "u", "-Syu"},
	Short:   "Update packages in the Arch Linux repository",
	Args:    cobra.MinimumNArgs(0),
	RunE: func(c *cobra.Command, args []string) error {
		pacArgs := append([]string{"-Syu"}, args...)
		return run.SudoStdTerminal(pacArgs...)
	},
}

var uninstallCmd = &cobra.Command{
	Use:     "uninstall <package(s)>",
	Aliases: []string{"-Rns", "rm", "remove", "delete"},
	Short:   "Delete an existing package from your system",
	RunE: func(c *cobra.Command, args []string) error {
		pacArgs := append([]string{"-Rns"}, args...)
		return run.SudoStdTerminal(pacArgs...)
	},
}

func init() {
	rootCmd.AddCommand(installCmd, updateCmd, uninstallCmd)
}
