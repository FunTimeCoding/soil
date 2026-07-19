package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue/issue_enricher"
	"testing"
)

func TestEnricher(t *testing.T) {
	assert.NotNil(
		t,
		issue_enricher.New(
			issue_enricher.WithConcernFunction(
				func(i *issue.Issue) []string {
					return []string{}
				},
			),
			issue_enricher.WithScoreFunction(
				func(*issue.Issue) float64 {
					return 0
				},
			),
			issue_enricher.WithCommentNameFilter([]string{}),
		),
	)
}
