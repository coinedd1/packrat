package cmd

import (
	"fmt"

	run "github.com/coinedd1/packrat/internal"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install <package(s)>",
	Aliases: []string{"i", "in", "ins", "S"},
	Short:   "Install a package from the Arch Linux repository",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		pacArgs := append([]string{"-S"}, args...)
		return run.SudoStdTerminal(pacArgs...)
	},
}

var updateCmd = &cobra.Command{
	Use:     "update <package(s)>",
	Aliases: []string{"upd", "u", "Syu"},
	Short:   "Update packages in the Arch Linux repository",
	Args:    cobra.MinimumNArgs(0),
	RunE: func(c *cobra.Command, args []string) error {
		pacArgs := append([]string{"-Syu"}, args...)
		return run.SudoStdTerminal(pacArgs...)
	},
}

var uninstallCmd = &cobra.Command{
	Use:     "uninstall <package(s)>",
	Aliases: []string{"Rns", "rm", "remove", "delete"},
	Short:   "Delete an existing package from your system",
	RunE: func(c *cobra.Command, args []string) error {
		pacArgs := append([]string{"-Rns"}, args...)
		return run.SudoStdTerminal(pacArgs...)
	},
}

var queryCmd = &cobra.Command{
	Use:     "query <keyword>",
	Aliases: []string{"Q", "q", "qu"},
	Short:   "Query pacman and search through existing packages on your system",
	RunE: func(c *cobra.Command, args []string) error {
		pacArgs := append([]string{"-Q"}, args...)
		out, err := run.CmdPrint(pacArgs)
		if err != nil {
			fmt.Println("query: package", args, "not found")
			return err
		}
		if out != "" {
			fmt.Println(out)
		}
		return nil
	},
}

var searchCmd = &cobra.Command{
	Use:     "search <keyword>",
	Aliases: []string{"Si", "srch", "lookup"},
	Short:   "Search the arch repos for information on a specific package",
	RunE: func(c *cobra.Command, args []string) error {
		pacArgs := append([]string{"-Si"}, args...)
		out, err := run.CmdPrint(pacArgs)
		if err != nil {
			fmt.Println("search: package", args, "not found")
			return err
		}
		if out != "" {
			fmt.Println(out)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd, updateCmd, uninstallCmd, queryCmd, searchCmd)
}

var cleanCmd = &cobra.Command{
	Use:     "clean",
	Aliases: []string{"Sc", "clear-cache"},
	Short:   "Delete cached files for unused packages",
	Args:    cobra.MaximumNArgs(0),
	RunE: func(c *cobra.Command, args []string) error {
		pacArgs := append([]string{"-Sc"}, args...)
		return run.SudoStdTerminal(pacArgs...)
	},
}

var cleanAllCmd = &cobra.Command{
	Use:     "all",
	Aliases: []string{"Scc", "clear-cache-all"},
	Short:   "Delete cached files for used AND unused packages",
	Args:    cobra.MaximumNArgs(0),
	RunE: func(c *cobra.Command, args []string) error {
		pacArgs := append([]string{"-Scc"}, args...)
		return run.SudoStdTerminal(pacArgs...)
	},
}

func init() {
	cleanCmd.AddCommand(cleanAllCmd)
	rootCmd.AddCommand(cleanCmd)
}
