package main

import (
	confluence "github.com/funtimecoding/go-library/pkg/atlassian/confluence/example"
	jira "github.com/funtimecoding/go-library/pkg/atlassian/jira/example"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/example/search"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/example/token_check"
	opsgenie "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/example"
)

func main() {
	jira.Changelog()

	if false {
		confluence.Overview()
		confluence.Export()
		confluence.Page()
		confluence.Search()
		confluence.Watch()
		confluence.Label()
		confluence.Search()
		confluence.Space()
		confluence.User()
		jira.Customer()
		jira.Issue()
		search.Search()
		jira.Watch()
		jira.CustomValue()
		token_check.TokenCheck()
		opsgenie.AddResponder()
		opsgenie.Alert()
	}
}
