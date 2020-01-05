package exporter

import (
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/oke-py/contributions/pkg/github"
)

func TestWriteMarkdown(t *testing.T) {
	c := []github.AggregatedContributionsCollection{
		{
			Repository:             "GoogleCloudPlatform/istio-samples",
			CommitCount:            0,
			IssueCount:             0,
			PullRequestCount:       2,
			PullRequestReviewCount: 0,
		},
		{
			Repository:             "GoogleContainerTools/skaffold",
			CommitCount:            15,
			IssueCount:             7,
			PullRequestCount:       8,
			PullRequestReviewCount: 0,
		},
		{
			Repository:             "kubernetes/website",
			CommitCount:            17,
			IssueCount:             28,
			PullRequestCount:       42,
			PullRequestReviewCount: 24,
		},
	}

	expected := heredoc.Doc(`
		| Repository                     | Commits | Issues  | PRs     | Reviews |
		| ------------------------------ | ------: | ------: | ------: | ------: |
		| GoogleCloudPlatform/istio-samples |       0 |       0 |       2 |       0 |
		| GoogleContainerTools/skaffold  |      15 |       7 |       8 |       0 |
		| kubernetes/website             |      17 |      28 |      42 |      24 |
	`)

	actual := WriteMarkdown(c)
	if actual != expected {
		t.Fatal("WriteMarkdown() does not render markdown correctly")
	}
}
