package example

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func WatchedState() {
	console.Format(
		"%s set: %v\n",
		constant.JiraWatchedIssuesEnvironment,
		environment.Exists(constant.JiraWatchedIssuesEnvironment),
	)
	c := jira.NewEnvironment()
	o := c.MustIssueOption()
	console.Format(
		"Environment client: loaded %v, keys %d\n",
		o.WatchedLoaded,
		len(o.WatchedIssues),
	)
	w := jira.NewEnvironment(jira.WithWatchedIssues())
	p := w.MustIssueOption()
	console.Format(
		"Explicit option: loaded %v, keys %d\n",
		p.WatchedLoaded,
		len(p.WatchedIssues),
	)

	for _, k := range p.WatchedIssues[:min(5, len(p.WatchedIssues))] {
		console.Format("  %s\n", k)
	}
}
