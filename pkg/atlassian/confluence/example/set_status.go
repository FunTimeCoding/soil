package example

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func SetStatus() {
	if false {
		c := confluence.NewEnvironment()
		result, e := c.SetPageStatus("6717441", constant.ConfluenceDraftStatus)

		if e != nil {
			console.Format("error: %v\n", e)

			return
		}

		console.Format(
			"status=%s version=%d\n",
			result.Raw.Status,
			result.Raw.Version.Number,
		)
	}
}
