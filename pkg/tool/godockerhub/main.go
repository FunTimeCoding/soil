package godockerhub

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/docker/hub"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/godockerhub/constant"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.String(
		argument.Image,
		"",
		"Image to list tags for (e.g. library/golang)",
	)
	a.Parse(version, gitHash, buildDate)
	image := a.GetString(argument.Image)

	if image == "" {
		fmt.Println("--image is required")

		return
	}

	c := hub.New()
	tags := c.Tags(image)
	limit := len(tags)

	if limit > maxDisplay {
		limit = maxDisplay
	}

	for _, t := range tags[:limit] {
		updated := ""

		if t.LastUpdated != nil {
			updated = t.LastUpdated.Format(time.DateMinute)
		}

		fmt.Printf("%-40s %s\n", t.Name, updated)
	}
}
