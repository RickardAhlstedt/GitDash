package repo

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

type RepoStatus struct {
	Path        string
	Name        string
	Branch      string
	Ahead       int
	Behind      int
	Dirty       bool
	StatusLines []string
}

func GetRepoStatus(repoPath string) (*RepoStatus, error) {
	name := filepath.Base(repoPath)

	branch, err := gitCommand(repoPath, "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return nil, fmt.Errorf("failed to get branch for %s: %v", name, err)
	}
	branch = strings.TrimSpace(branch)

	remoteInfo, _ := gitCommand(repoPath, "rev-list", "--left-right", "--count", "HEAD...@{u}")
	ahead, behind := 0, 0
	if remoteInfo != "" {
		parts := strings.Fields(remoteInfo)
		if len(parts) == 2 {
			if _, err := fmt.Sscanf(parts[0], "%d", &ahead); err != nil {
				return nil, fmt.Errorf("failed to parse ahead count for %s: %v", name, err)
			}
			if _, err := fmt.Sscanf(parts[1], "%d", &behind); err != nil {
				return nil, fmt.Errorf("failed to parse behind count for %s: %v", name, err)
			}
		}
	}

	statusOutput, _ := gitCommand(repoPath, "status", "--porcelain")
	dirty := strings.TrimSpace(statusOutput) != ""

	lines := []string{}
	if dirty {
		for _, line := range strings.Split(statusOutput, "\n") {
			if strings.TrimSpace(line) != "" {
				lines = append(lines, line)
			}
		}
	}

	return &RepoStatus{
		Path:        repoPath,
		Name:        name,
		Branch:      branch,
		Ahead:       ahead,
		Behind:      behind,
		Dirty:       dirty,
		StatusLines: lines,
	}, nil
}

func gitCommand(repoPath string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = repoPath
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}
