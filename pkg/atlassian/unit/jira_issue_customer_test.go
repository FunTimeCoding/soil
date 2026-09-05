package unit

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/soil/pkg/assert"
	atlassian "github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue/customer"
	strings "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestIssue(t *testing.T) {
	actual := customer.New(
		&models.CustomerRequestScheme{
			IssueKey: strings.UpperAlfa,
			RequestFieldValues: []*models.CustomerRequestRequestFieldValueScheme{
				{
					FieldID: atlassian.JiraSummaryField,
					Value:   strings.UpperBravo,
				},
				{
					FieldID: atlassian.JiraDescriptionField,
					Value:   strings.UpperCharlie,
				},
			},
			CurrentStatus: &models.CustomerRequestCurrentStatusScheme{
				Status: atlassian.ServiceDeskResolved,
			},
			Links: &models.CustomerRequestLinksScheme{Web: strings.UpperDelta},
		},
	)
	actual.Raw = nil
	assert.Any(
		t,
		&customer.Issue{
			Key:    "Alfa",
			Status: "Resolved",
			Title:  "Bravo",
			Body:   "Charlie",
			Link:   "Delta",
		},
		actual,
	)
}
