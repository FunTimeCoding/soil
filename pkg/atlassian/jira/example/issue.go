package example

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/davecgh/go-spew/spew"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Issue() {
	j := common.Jira()
	p := j.DefaultProjectKey()
	issueType := issue.TaskType
	summary := "Stub summary"
	description := "Stub description"
	f := constant.Format.Copy()
	var i *jira.Issue

	if true {
		i = j.MustNewIssue(p, issueType, summary, description)
	}

	if false {
		i = j.MustNewIssueUnverified(p, issueType, summary, description)
	}

	fmt.Println("Prepared:")
	spew.Dump(i)

	if false {
		fmt.Println("Created:")
		fmt.Println(j.MustCreateNative(i).Format(f))
	}
}
