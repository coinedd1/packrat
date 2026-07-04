package run

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func SudoStdTerminal(line ...string) error {
	full := append([]string{"pacman"}, line...)
	cmd := exec.Command("sudo", full...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func CmdPrint(line ...string) (string, error) {
	out, err := exec.Command("pacman", line...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("pacman: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}
