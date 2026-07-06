package status

import (
	"github.com/funtimecoding/soil/pkg/expression"
	"strings"
)

func Parse(s string) *Result {
	return &Result{
		Loaded: strings.TrimSpace(
			expression.SubMatch(`Loaded: (.+)`, s),
		),
		Active: strings.TrimSpace(
			expression.SubMatch(`Active: (.+) since (.+);`, s),
		),
		MainProcess: strings.TrimSpace(
			expression.SubMatch(`Main PID: (\d+)`, s),
		),
		Memory: strings.TrimSpace(
			expression.SubMatch(`Memory: (.+)`, s),
		),
		Processor: strings.TrimSpace(
			expression.SubMatch(`CPU: (.+)`, s),
		),
	}
}
