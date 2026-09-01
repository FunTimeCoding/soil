package gosentry

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
	"github.com/funtimecoding/soil/pkg/system/environment"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
)

func showIssue(shortIdentifier string) {
	c := sentry.NewEnvironment()
	o := environment.Required(constant.OrganizationEnvironment)
	i := c.MustIssueByShortIdentifier(o, shortIdentifier)

	if i == nil {
		console.Format("Issue not found: %s\n", shortIdentifier)

		return
	}

	r := i.Raw
	console.Format("Issue:    %s\n", r.ShortIdentifier)
	console.Format("Title:    %s\n", r.Title)
	console.Format("Project:  %s\n", r.Project.Name)
	console.Format("Link:     %s\n", r.Permalink)
	console.Format("Status:   %s\n", r.Status)
	console.Format("Level:    %s\n", r.Level)
	console.Format("Events:   %s\n", r.Count)

	if r.FirstSeen != nil {
		console.Format(
			"First:    %s\n",
			r.FirstSeen.Format(timeConstant.DateMinute),
		)
	}

	if r.LastSeen != nil {
		console.Format(
			"Last:     %s\n",
			r.LastSeen.Format(timeConstant.DateMinute),
		)
	}

	if r.Culprit != "" {
		console.Format("Culprit:  %s\n", r.Culprit)
	}

	e := c.MustLatestEvent(o, r.Identifier)
	console.Line()
	console.Format("Latest Event: %s\n", e.EventIdentifier)

	if e.DateCreated != nil {
		console.Format(
			"Date:     %s\n",
			e.DateCreated.Format(timeConstant.DateMinute),
		)
	}

	if len(e.Entries) > 0 {
		console.Line()
		printEventEntries(e.Entries)
	}
}
