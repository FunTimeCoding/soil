package unit

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue/customer"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestIssue(t *testing.T) {
	actual := customer.New(
		&models.CustomerRequestScheme{
			IssueKey: stringsConstant.UpperAlfa,
			RequestFieldValues: []*models.CustomerRequestRequestFieldValueScheme{
				{
					FieldID: constant.JiraSummaryField,
					Value:   stringsConstant.UpperBravo,
				},
				{
					FieldID: constant.JiraDescriptionField,
					Value:   stringsConstant.UpperCharlie,
				},
			},
			CurrentStatus: &models.CustomerRequestCurrentStatusScheme{
				Status: constant.ServiceDeskResolved,
			},
			Links: &models.CustomerRequestLinksScheme{
				Web: stringsConstant.UpperDelta,
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
