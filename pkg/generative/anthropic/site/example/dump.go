package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/site"
)

func Dump() {
	s := site.New()
	r := s.ReadUsage()

	if r == nil {
		console.Line("no usage data")

		return
	}

	console.Format(
		"Session %d%%  resets %s\n",
		r.SessionPercent,
		r.SessionReset,
	)
	console.Format(
		"Weekly  %d%%  resets %s\n",
		r.WeeklyAllPercent,
		r.WeeklyAllReset,
	)
	console.Format("Fable   %d%%  resets %s\n", r.FablePercent, r.FableReset)
}
