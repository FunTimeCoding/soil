package main

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/example"
	jira "github.com/funtimecoding/soil/pkg/atlassian/jira/example"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/example/search"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/example/token_check"
	opsgenie "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/example"
)

func main() {
	jira.WatchedState()

	if false {
		jira.Changelog()
		example.Overview()
		example.Export()
		example.Page()
		example.Search()
		example.Watch()
		example.Label()
		example.Search()
		example.Space()
		example.User()
		jira.Customer()
		jira.Issue()
		search.Search()
		jira.Watch()
		jira.WatchedState()
		jira.CustomValue()
		token_check.TokenCheck()
		opsgenie.AddResponder()
		opsgenie.Alert()
	}
}
