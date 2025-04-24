package repo

import (
	"bytes"
	"fmt"
	"os/exec"
)

func FetchOrigin(path string, quiet bool) error {
	var cmd *exec.Cmd
	if quiet {
		cmd = exec.Command("git", "-C", path, "fetch", "--all", "--quiet")
	} else {
		cmd = exec.Command("git", "-C", path, "fetch", "--all")
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git fetch failed for %s: %s", path, stderr.String())
	}
	return nil
}
