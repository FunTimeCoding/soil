package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/common"
	"time"
)

func Changelog() {
	k := environment.Required(constant.DefaultProjectKeyEnvironment)
	j := common.Jira()
	fmt.Println("Search (with changelog)...")
	start := time.Now()
	issues := j.MustSearch(
		"project = %s AND status NOT IN (Backlog, Closed) ORDER BY updated ASC",
		k,
	)
	elapsed := time.Since(start)
	fmt.Printf("Fetched %d issues in %s\n\n", len(issues), elapsed)
	fmt.Printf(
		"  %-10s %-22s %6s %6s\n",
		"KEY",
		"STATUS",
		"CHANGE",
		"TRANS",
	)

	for _, i := range issues {
		changeAge := int(time.Since(i.ChangeTime()).Hours() / 24)
		transStr := "n/a"

		if i.Raw.Changelog != nil {
			var lastTransitionTime time.Time

			for _, h := range i.Raw.Changelog.Histories {
				for _, item := range h.Items {
					if item.Field == "status" {
						t, f := time.Parse(
							constant.TimeFormat,
							h.Created,
						)

						if f != nil {
							continue
						}

						if t.After(lastTransitionTime) {
							lastTransitionTime = t
						}
					}
				}
			}

			if !lastTransitionTime.IsZero() {
				transStr = fmt.Sprintf(
					"%dd",
					int(time.Since(lastTransitionTime).Hours()/24),
				)
			}
		}

		fmt.Printf(
			"  %-10s %-22s %5dd %6s\n",
			i.Key,
			i.Status,
			changeAge,
			transStr,
		)
	}
}
