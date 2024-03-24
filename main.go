package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/cli/go-gh"
)

type RepoStats struct {
	FullName        string `json:"full_name"`
	StargazersCount int    `json:"stargazers_count"`
	ForksCount      int    `json:"forks_count"`
	OpenIssuesCount int    `json:"open_issues_count"`
	WatchersCount   int    `json:"watchers_count"`
	DefaultBranch   string `json:"default_branch"`
	Archived        bool   `json:"archived"`
}

func main() {
	fmt.Printf("Debug: Received args: %v\n", os.Args)

	var args []string
	var owner, repo string

	if len(os.Args) == 2 {
		args = strings.Split(os.Args[1], " ")
	} else {
		args = os.Args[1:] // Skip the program name
	}

	for i, arg := range args {
		if arg == "-owner" && i+1 < len(args) {
			owner = args[i+1]
		} else if arg == "-repo" && i+1 < len(args) {
			repo = args[i+1]
		}
	}

	if owner == "" || repo == "" {
		fmt.Println("Usage: gh repo-stats -owner <owner> -repo <repo>")
		os.Exit(1)
	}

	ghArgs := []string{"api", fmt.Sprintf("repos/%s/%s", owner, repo)}
	stdOut, _, err := gh.Exec(ghArgs...)
	if err != nil {
		fmt.Printf("Error fetching repository data: %s\n", err)
		return
	}

	var stats RepoStats
	if err := json.Unmarshal(stdOut.Bytes(), &stats); err != nil {
		fmt.Printf("Error parsing JSON: %s\n", err)
		return
	}

	fmt.Printf("Repository: %s\n🌟 Stars: %d\n🍴 Forks: %d\n🔓 Open Issues: %d\n👀 Watchers: %d\n🔖 Default Branch: %s\n📦 Archived: %t\n",
		stats.FullName, stats.StargazersCount, stats.ForksCount, stats.OpenIssuesCount, stats.WatchersCount, stats.DefaultBranch, stats.Archived)
}
