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
		step.New("alfa.example", constant.Refused),
		step.New("alfa.example", constant.Refused),
		step.New("bravo.example", constant.NoRoute),
		step.New("alfa.example", constant.TimedOut),
		step.New("bravo.example", constant.NoRoute),
		step.New("alfa.example", ""),
		step.New("alfa.example", ""),
		step.New("bravo.example", ""),
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
