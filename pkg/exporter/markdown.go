package exporter

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/oke-py/contributions/pkg/github"
)

// WriteMarkdown convert array of contributions collection to markdown.
func WriteMarkdown(c []github.AggregatedContributionsCollection) string {
	header := heredoc.Doc(`
		| Repository                     | Commits | Issues  | PRs     | Reviews |
		| ------------------------------ | ------: | ------: | ------: | ------: |
	`)
	md := fmt.Sprint(header)

	for _, v := range c {
		md += fmt.Sprintf(
			"| %-30s | %7d | %7d | %7d | %7d |\n",
			v.Repository,
			v.CommitCount,
			v.IssueCount,
			v.PullRequestCount,
			v.PullRequestReviewCount,
		)
	}

	return md
}
