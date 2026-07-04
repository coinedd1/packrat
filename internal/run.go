package run

import (
	"fmt"
	"os/exec"
	"strings"
)

func Command(line string) (string, error) {
	args := []string{}
	args = append(args, "pacman", "-S", line)
	out, err := exec.Command("sudo", args...).CombinedOutput()
	fmt.Println(string(out))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
