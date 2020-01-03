package exporter

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/oke-py/contributions/pkg/github"
)

// WriteMarkdown convert array of contributions collection to markdown.
func WriteMarkdown(c map[string]github.AggregatedContributionsCollection) {
	header := heredoc.Doc(`
		| Repository                     | Commits | Issues  | PRs     | Reviews |
		| ------------------------------ | ------: | ------: | ------: | ------: |
	`)
	fmt.Print(header)

	for repo, v := range c {
		fmt.Fprintf(
			os.Stdout,
			"| %-30s | %7d | %7d | %7d | %7d |\n",
			repo,
			v.CommitCount,
			v.IssueCount,
			v.PullRequestCount,
			v.PullRequestReviewCount,
		)
	}
}
