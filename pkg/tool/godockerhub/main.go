package godockerhub

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/docker/hub"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/tool/godockerhub/constant"
)

func Main() {
	r := reporter.New(constant.Identity.Name()).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.String(
		argumentConstant.Image,
		"",
		"Image to list tags for (e.g. library/golang)",
	)
	a.Parse()
	image := a.GetString(argumentConstant.Image)

	if image == "" {
		console.Line("--image is required")

		return
	}

	c := hub.New()
	tags := c.MustTags(image)
	limit := len(tags)

	if limit > constant.MaxDisplay {
		limit = constant.MaxDisplay
	}

	for _, t := range tags[:limit] {
		updated := ""

		if t.LastUpdated != nil {
			updated = t.LastUpdated.Format(timeConstant.DateMinute)
		}

		console.Format("%-40s %s\n", t.Name, updated)
	}
}
