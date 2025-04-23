package repo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
)

func FindGitRepos(rootPaths []string, ignorePatterns []string) ([]string, error) {
	var repos []string

	for _, root := range rootPaths {
		if err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err // Propagate the error
			}

			// Ignore match
			for _, pattern := range ignorePatterns {
				if matched, _ := doublestar.Match(pattern, path); matched {
					if d.IsDir() {
						return filepath.SkipDir
					}
					return nil
				}
			}

			// Git repo detection
			if d.IsDir() && d.Name() == ".git" {
				repoRoot := filepath.Dir(path)
				repos = append(repos, repoRoot)
				return filepath.SkipDir // don't walk inside .git
			}

			return nil
		}); err != nil {
			return nil, fmt.Errorf("error walking directory: %v", err)
		}
	}

	return repos, nil
}
