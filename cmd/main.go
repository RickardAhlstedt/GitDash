package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/RickardAhlstedt/GitDash/config"
	"github.com/RickardAhlstedt/GitDash/repo"
)

func main() {
	configPath := flag.String("config", "", "Path to gitdash config file")
	// noFetch := flag.Bool("no-fetch", false, "Don't run git fetch (future)")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	fmt.Println("📂 Scanning paths for git repositories...")
	repos, err := repo.FindGitRepos(cfg.Paths, cfg.Ignore)
	if err != nil {
		log.Fatalf("❌ Failed to find repos: %v", err)
	}
	if len(repos) == 0 {
		fmt.Println("⚠️  No git repositories found.")
		os.Exit(0)
	}

	fmt.Printf("🔍 Found %d repositories\n\n", len(repos))

	for _, path := range repos {
		status, err := repo.GetRepoStatus(path)
		if err != nil {
			fmt.Printf("❌ %s - error: %v\n", path, err)
			continue
		}

		dirty := ""
		if status.Dirty {
			dirty = "✴"
		}

		fmt.Printf("📁 %-30s [%s] ↑%d ↓%d %s\n",
			status.Name, status.Branch, status.Ahead, status.Behind, dirty)
	}
}
