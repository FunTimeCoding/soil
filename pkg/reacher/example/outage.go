package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/reacher"
	"github.com/funtimecoding/soil/pkg/reacher/example/step"
)

func Outage() {
	r := reacher.New()
	script := []*step.Step{
		step.New("sentry2.rz.adition.net", constant.Refused),
		step.New("sentry2.rz.adition.net", constant.Refused),
		step.New("netbox.v10s.net", constant.NoRoute),
		step.New("sentry2.rz.adition.net", constant.TimedOut),
		step.New("netbox.v10s.net", constant.NoRoute),
		step.New("sentry2.rz.adition.net", ""),
		step.New("sentry2.rz.adition.net", ""),
		step.New("netbox.v10s.net", ""),
	}

	for _, s := range script {
		e := r.Succeed(s.Host)

		if s.Reason != "" {
			e = r.Fail(s.Host, s.Reason)
		}

		if e != nil {
			fmt.Println(e)
		}
	}
}
