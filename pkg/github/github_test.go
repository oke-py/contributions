package github

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	c := ContributionsCollection{
		CommitContributionsByRepository: []commitContributions{
			{
				Contributions: contributions{
					TotalCount: 2,
				},
				Repository: repository{
					NameWithOwner: "kubernetes/website",
				},
			},
		},
		IssueContributionsByRepository: []issueContributions{
			{
				Contributions: contributions{
					TotalCount: 1,
				},
				Repository: repository{
					NameWithOwner: "kubernetes/kubernetes",
				},
			},
			{
				Contributions: contributions{
					TotalCount: 3,
				},
				Repository: repository{
					NameWithOwner: "kubernetes/website",
				},
			},
		},
		PullRequestContributionsByRepository: []pullRequestContributions{
			{
				Contributions: contributions{
					TotalCount: 1,
				},
				Repository: repository{
					NameWithOwner: "kubernetes/kubernetes",
				},
			},
			{
				Contributions: contributions{
					TotalCount: 5,
				},
				Repository: repository{
					NameWithOwner: "kubernetes/website",
				},
			},
		},
		PullRequestReviewContributionsByRepository: []pullRequestReviewContributions{
			{
				Contributions: contributions{
					TotalCount: 5,
				},
				Repository: repository{
					NameWithOwner: "kubernetes/website",
				},
			},
		},
	}

	expected := map[string]AggregatedContributionsCollection{
		"kubernetes/website": {
			CommitCount:            2,
			IssueCount:             3,
			PullRequestCount:       5,
			PullRequestReviewCount: 5,
		},
		"kubernetes/kubernetes": {
			CommitCount:            0,
			IssueCount:             1,
			PullRequestCount:       1,
			PullRequestReviewCount: 0,
		},
	}

	actual := c.Convert()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatal("Convert() does not convert struct correctly")
	}
}
