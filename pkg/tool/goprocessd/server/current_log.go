package server

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/process"
)

func currentLog(p *process.Process, name string) string {
	current, older := p.CurrentLog()

	if len(current) == 0 && older == 0 {
		return "ok"
	}

	if older == 0 {
		return lines(current)
	}

	if len(current) == 0 {
		return fmt.Sprintf("(%d older lines, use: log %s all)", older, name)
	}

	return fmt.Sprintf(
		"%s\n(%d older lines, use: log %s all)",
		join.NewLine(current),
		older,
		name,
	)
}
