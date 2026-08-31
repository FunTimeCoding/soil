package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/site"
)

func Dump() {
	s := site.New()
	r := s.ReadUsage()

	if r == nil {
		fmt.Println("no usage data")

		return
	}

	fmt.Printf("Session %d%%  resets %s\n", r.SessionPercent, r.SessionReset)
	fmt.Printf(
		"Weekly  %d%%  resets %s\n",
		r.WeeklyAllPercent,
		r.WeeklyAllReset,
	)
	fmt.Printf("Fable   %d%%  resets %s\n", r.FablePercent, r.FableReset)
}
