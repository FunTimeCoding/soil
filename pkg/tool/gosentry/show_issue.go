package gosentry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
	"github.com/funtimecoding/soil/pkg/errors/sentry/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/time"
)

func showIssue(shortIdentifier string) {
	c := sentry.NewEnvironment()
	o := environment.Required(constant.OrganizationEnvironment)
	i := c.MustIssueByShortIdentifier(o, shortIdentifier)

	if i == nil {
		fmt.Printf("Issue not found: %s\n", shortIdentifier)

		return
	}

	r := i.Raw
	fmt.Printf("Issue:    %s\n", r.ShortID)
	fmt.Printf("Title:    %s\n", r.Title)
	fmt.Printf("Project:  %s\n", r.Project.Name)
	fmt.Printf("Link:     %s\n", r.Permalink)
	fmt.Printf("Status:   %s\n", r.Status)
	fmt.Printf("Level:    %s\n", r.Level)
	fmt.Printf("Events:   %s\n", r.Count)

	if r.FirstSeen != nil {
		fmt.Printf(
			"First:    %s\n",
			r.FirstSeen.Format(time.DateMinute),
		)
	}

	if r.LastSeen != nil {
		fmt.Printf(
			"Last:     %s\n",
			r.LastSeen.Format(time.DateMinute),
		)
	}

	if r.Culprit != "" {
		fmt.Printf("Culprit:  %s\n", r.Culprit)
	}

	e := c.MustLatestEvent(o, r.ID)
	fmt.Println()
	fmt.Printf("Latest Event: %s\n", e.EventID)

	if e.DateCreated != nil {
		fmt.Printf(
			"Date:     %s\n",
			e.DateCreated.Format(time.DateMinute),
		)
	}

	if len(e.Entries) > 0 {
		fmt.Println()
		printEventEntries(e.Entries)
	}
}
