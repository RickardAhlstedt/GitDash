package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/RickardAhlstedt/GitDash/config"
	"github.com/RickardAhlstedt/GitDash/repo"
	"github.com/RickardAhlstedt/GitDash/style"
	"github.com/RickardAhlstedt/GitDash/tui"
)

func main() {
	configPath := flag.String("config", "", "Path to gitdash config file")
	sortByFlag := flag.String("sort", "", "Sort by: name | branch | ahead | behind")
	tuiFlag := flag.Bool("tui", false, "Use TUI, false by default")
	themeFlag := flag.String("theme", "", "Override theme (e.g. dracula, nord, lolcat)")
	fetchFlag := flag.Bool("fetch", false, "Fetch origin before status")
	quietFlag := flag.Bool("quiet", false, "Suppress any output from fetching the origin")

	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}
	if *themeFlag != "" {
		cfg.Theme.Name = *themeFlag
	}
	style.SetTheme(cfg.Theme.Name, cfg.Theme.Colors)

	sortBy := strings.ToLower(*sortByFlag)
	if sortBy == "" {
		sortBy = strings.ToLower(cfg.SortBy)
	}
	if sortBy == "" {
		sortBy = "name"
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

	if cfg.Fetch || *fetchFlag {
		fmt.Println("🔄 Fetching origin for all repositories...")
		for _, path := range repos {
			if err := repo.FetchOrigin(path, *quietFlag); err != nil {
				fmt.Printf("⚠️  Failed to fetch for %s: %v\n", path, err)
			}
		}
		fmt.Println()
	}

	var statuses []*repo.RepoStatus
	for _, path := range repos {
		status, err := repo.GetRepoStatus(path)
		if err == nil {
			statuses = append(statuses, status)
		}
	}

	switch sortBy {
	case "branch":
		sort.Slice(statuses, func(i, j int) bool {
			return statuses[i].Branch < statuses[j].Branch
		})
	case "ahead":
		sort.Slice(statuses, func(i, j int) bool {
			return statuses[i].Ahead > statuses[j].Ahead
		})
	case "behind":
		sort.Slice(statuses, func(i, j int) bool {
			return statuses[i].Behind > statuses[j].Behind
		})
	default:
		sort.Slice(statuses, func(i, j int) bool {
			return statuses[i].Name < statuses[j].Name
		})
	}

	if *tuiFlag {
		tui.Render(statuses, sortBy)
	} else {
		for _, s := range statuses {
			dirty := ""
			if s.Dirty {
				dirty = "✴"
			}
			fmt.Printf("📁 %-30s [%s] ↑%d ↓%d %s\n", s.Name, s.Branch, s.Ahead, s.Behind, dirty)
		}
	}

}
