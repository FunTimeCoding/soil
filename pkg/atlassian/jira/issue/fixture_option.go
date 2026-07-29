package issue

import (
	"github.com/andygrunwald/go-jira"
	atlassian "github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/field_map"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue/option"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func FixtureOption() *option.Issue {
	return option.New(
		locator.New(webConstant.Localhost).Insecure().String(),
		"test",
		[]string{},
		[]string{atlassian.JiraClosed},
		field_map.New([]jira.Field{}),
	)
}
