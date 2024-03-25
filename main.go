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

func getContributorsFromPage(owner, repo string, pageNum int) []interface{} {
	ghArgs := []string{"api", fmt.Sprintf("repos/%s/%s/contributors?page=%d", owner, repo, pageNum)}
	stdOut, _, err := gh.Exec(ghArgs...)
	if err != nil {
		fmt.Printf("Error fetching contributors: %s\n", err)
		return nil
	}

	var contributors []interface{}
	if err := json.Unmarshal(stdOut.Bytes(), &contributors); err != nil {
		fmt.Printf("Error parsing JSON: %s\n", err)
		return nil
	}

	return contributors
}

func getCommitsFromPage(owner, repo string, pageNum int) []interface{} {
	ghArgs := []string{"api", fmt.Sprintf("repos/%s/%s/commits?page=%d", owner, repo, pageNum)}
	stdOut, _, err := gh.Exec(ghArgs...)
	if err != nil {
		fmt.Printf("Error fetching commits: %s\n", err)
		return nil
	}

	var commits []interface{}
	if err := json.Unmarshal(stdOut.Bytes(), &commits); err != nil {
		fmt.Printf("Error parsing JSON: %s\n", err)
		return nil
	}

	return commits
}

func fetchContributorsCount(owner, repo string) int {
	pageNum := 1
	totalContributors := 0

	for {
		contributors := getContributorsFromPage(owner, repo, pageNum)
		if len(contributors) == 0 {
			break
		}
		totalContributors += len(contributors)
		pageNum++
	}

	return totalContributors
}

func fetchTotalCommits(owner, repo string) int {
	pageNum := 1
	totalCommits := 0

	for {
		commits := getCommitsFromPage(owner, repo, pageNum)
		if len(commits) == 0 {
			break
		}
		totalCommits += len(commits)
		pageNum++
	}

	return totalCommits
}

func main() {
	fmt.Printf("Debug: Received args: %v\n", os.Args)

	var args []string
	var owner, repo string

	if len(os.Args) == 2 {
		args = strings.Split(os.Args[1], " ")
	} else {
		args = os.Args[1:]
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

	fmt.Printf("Debug: Owner: %s, Repo: %s\n", owner, repo)

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

	contributorsCount := fetchContributorsCount(owner, repo)
	totalCommits := fetchTotalCommits(owner, repo)

	fmt.Printf("Repository: %s\n", stats.FullName)
	fmt.Printf("🌟 Stars: %d\n", stats.StargazersCount)
	fmt.Printf("🍴 Forks: %d\n", stats.ForksCount)
	fmt.Printf("🔓 Open Issues: %d\n", stats.OpenIssuesCount)
	fmt.Printf("👀 Watchers: %d\n", stats.WatchersCount)
	fmt.Printf("🔖 Default Branch: %s\n", stats.DefaultBranch)
	fmt.Printf("📦 Archived: %t\n", stats.Archived)
	fmt.Printf("👥 Contributors: %d\n", contributorsCount)
	fmt.Printf("📌 Total Commits: %d\n", totalCommits)

}
