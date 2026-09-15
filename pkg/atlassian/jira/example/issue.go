package example

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Issue() {
	j := common.Jira()
	p := j.DefaultProjectKey()
	issueType := constant.JiraTaskType
	summary := "Stub summary"
	description := "Stub description"
	f := constant.JiraFormat.Copy()
	var i *jira.Issue

	if true {
		i = j.MustNewIssue(p, issueType, summary, description)
	}

	if false {
		i = j.MustNewIssueUnverified(p, issueType, summary, description)
	}

	console.Line("Prepared:")
	console.Line(notation.MarshalIndent(i))

	if false {
		console.Line("Created:")
		console.Line(j.MustCreateNative(i).Format(f))
	}
}
