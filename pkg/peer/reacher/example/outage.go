package example

import (
	"fmt"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/peer/reacher"
	"github.com/funtimecoding/soil/pkg/peer/reacher/example/step"
)

func Outage() {
	r := reacher.New(memory.New())
	script := []*step.Step{
		step.New("alfa.example", errorsConstant.Refused),
		step.New("alfa.example", errorsConstant.Refused),
		step.New("bravo.example", errorsConstant.NoRoute),
		step.New("alfa.example", errorsConstant.TimedOut),
		step.New("bravo.example", errorsConstant.NoRoute),
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
