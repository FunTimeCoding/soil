package unit_test

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/search_result"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/space"
	confluenceUser "github.com/funtimecoding/soil/pkg/atlassian/confluence/user"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/field_map"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue/issue_enricher"
	issueOption "github.com/funtimecoding/soil/pkg/atlassian/jira/issue/option"
	alertOption "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert/option"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/override"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/rotation"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/schedule"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team_map"
	opsgenieUser "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/user"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/user_map"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	rawSchedule "github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
	rawTeam "github.com/opsgenie/opsgenie-go-sdk-v2/team"
	rawUser "github.com/opsgenie/opsgenie-go-sdk-v2/user"
	"testing"
)

func TestEntities(t *testing.T) {
	assert.NotNil(t, page.New(response.NewPage(), upper.Alfa))
	assert.NotNil(t, search_result.New(response.NewResult()))
	assert.NotNil(t, space.New(response.NewSpace()))
	assert.NotNil(t, confluenceUser.New(response.NewUser()))
	assert.NotNil(t, field_map.New([]jira.Field{}))
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
	assert.NotNil(
		t,
		issueOption.New(
			upper.Alfa,
			upper.Bravo,
			[]string{},
			[]string{constant.Done},
			nil,
		),
	)
	assert.NotNil(
		t,
		alertOption.New(
			nil,
			nil,
			upper.Alfa,
			nil,
			nil,
			nil,
			nil,
		),
	)
	assert.NotNil(t, override.New(&rawSchedule.ScheduleOverride{}))
	assert.NotNil(t, rotation.New(&rawSchedule.Rotation{}))
	assert.NotNil(t, schedule.New(&rawSchedule.Schedule{}, nil))
	assert.NotNil(t, team_map.New([]*team.Team{}))
	assert.NotNil(t, team.New(&rawTeam.ListedTeams{}))
	assert.NotNil(t, user_map.New([]*opsgenieUser.User{}))
	assert.NotNil(t, opsgenieUser.New(&rawUser.User{}))
}
