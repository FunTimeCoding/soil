package unit_test

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue/customer"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestIssue(t *testing.T) {
	actual := customer.New(
		&models.CustomerRequestScheme{
			IssueKey: upper.Alfa,
			RequestFieldValues: []*models.CustomerRequestRequestFieldValueScheme{
				{FieldID: customer.SummaryField, Value: upper.Bravo},
				{FieldID: customer.DescriptionField, Value: upper.Charlie},
			},
			CurrentStatus: &models.CustomerRequestCurrentStatusScheme{
				Status: constant.ServiceDeskResolved,
			},
			Links: &models.CustomerRequestLinksScheme{
				Web: upper.Delta,
			},
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
