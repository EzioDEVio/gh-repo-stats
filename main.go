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
}

func main() {
	fmt.Printf("Debug: Received args: %v\n", os.Args)

	var args []string
	var owner, repo string

	// If all arguments are passed as a single string, split them manually.
	if len(os.Args) == 2 {
		args = strings.Split(os.Args[1], " ")
	} else {
		args = os.Args[1:] // Skip the program name
	}

	// Parse -owner and -repo arguments
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

	fmt.Printf("Debug: Owner: %s, Repo: %s\n", owner, repo)

	// Construct the GitHub CLI command
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

	fmt.Printf("Repository: %s\nStars: %d\nForks: %d\nOpen Issues: %d\n",
		stats.FullName, stats.StargazersCount, stats.ForksCount, stats.OpenIssuesCount)
}
