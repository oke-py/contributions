package github

import (
	"context"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type contributions struct {
	TotalCount githubv4.Int
}

type repository struct {
	NameWithOwner githubv4.String
}

type commitContributions struct {
	Contributions contributions
	Repository    repository
}

type issueContributions struct {
	Contributions contributions
	Repository    repository
}

type pullRequestContributions struct {
	Contributions contributions
	Repository    repository
}

type pullRequestReviewContributions struct {
	Contributions contributions
	Repository    repository
}

// ContributionsCollection is a part of GraphQL response.
type ContributionsCollection struct {
	CommitContributionsByRepository            []commitContributions
	IssueContributionsByRepository             []issueContributions
	PullRequestContributionsByRepository       []pullRequestContributions
	PullRequestReviewContributionsByRepository []pullRequestReviewContributions
}

// AggregatedContributionsCollection has contribution count by type.
type AggregatedContributionsCollection struct {
	Repository             string
	CommitCount            int
	IssueCount             int
	PullRequestCount       int
	PullRequestReviewCount int
}

var query struct {
	User struct {
		ContributionsCollection ContributionsCollection `graphql:"contributionsCollection(from: $from, to: $to)"`
	} `graphql:"user(login: $user)"`
}

// GetContributions get GraphQL response from GitHub GraphQL API v4.
func GetContributions() ContributionsCollection {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)

	loc, _ := time.LoadLocation("Asia/Tokyo")
	variables := map[string]interface{}{
		"user": githubv4.String("oke-py"),
		"from": githubv4.DateTime{Time: time.Date(2019, 10, 1, 0, 0, 0, 0, loc)},
		"to":   githubv4.DateTime{Time: time.Date(2019, 11, 1, 0, 0, 0, 0, loc)},
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		// Handle error.
		fmt.Printf("error: %v", err)
	}

	return query.User.ContributionsCollection
}

// Convert transforms GraphQL response to map of AggregatedContributionsCollection.
func (c ContributionsCollection) Convert() []AggregatedContributionsCollection {
	m := make(map[string]AggregatedContributionsCollection)

	for _, v := range c.CommitContributionsByRepository {
		m[string(v.Repository.NameWithOwner)] = AggregatedContributionsCollection{
			Repository:  string(v.Repository.NameWithOwner),
			CommitCount: int(v.Contributions.TotalCount),
		}
	}

	for _, v := range c.IssueContributionsByRepository {
		if v2, ok := m[string(v.Repository.NameWithOwner)]; ok {
			v2.IssueCount = int(v.Contributions.TotalCount)
			m[string(v.Repository.NameWithOwner)] = v2
		} else {
			m[string(v.Repository.NameWithOwner)] = AggregatedContributionsCollection{
				Repository: string(v.Repository.NameWithOwner),
				IssueCount: int(v.Contributions.TotalCount),
			}
		}
	}

	for _, v := range c.PullRequestContributionsByRepository {
		if v2, ok := m[string(v.Repository.NameWithOwner)]; ok {
			v2.PullRequestCount = int(v.Contributions.TotalCount)
			m[string(v.Repository.NameWithOwner)] = v2
		} else {
			m[string(v.Repository.NameWithOwner)] = AggregatedContributionsCollection{
				Repository:       string(v.Repository.NameWithOwner),
				PullRequestCount: int(v.Contributions.TotalCount),
			}
		}
	}

	for _, v := range c.PullRequestReviewContributionsByRepository {
		if v2, ok := m[string(v.Repository.NameWithOwner)]; ok {
			v2.PullRequestReviewCount = int(v.Contributions.TotalCount)
			m[string(v.Repository.NameWithOwner)] = v2
		} else {
			m[string(v.Repository.NameWithOwner)] = AggregatedContributionsCollection{
				Repository:             string(v.Repository.NameWithOwner),
				PullRequestReviewCount: int(v.Contributions.TotalCount),
			}
		}
	}

	s := make([]AggregatedContributionsCollection, 0)

	for _, v := range m {
		s = append(s, v)
	}

	sort.SliceStable(s, func(i, j int) bool { return s[i].Repository < s[j].Repository })

	return s
}
