package cmd

import (
	"fmt"
	"strings"

	run "github.com/coinedd1/packrat/internal"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install <package(s)>",
	Aliases: []string{"in", "inst"},
	Short:   "Install a package from the Arch Linux repository",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		line := strings.Join(args, " ")
		out, err := run.Command(line)
		if err != nil {
			return err
		}
		if out != "" {
			fmt.Println(out)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
