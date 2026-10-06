package parser

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/section"
	"strings"
)

func Parse(body string) []*section.Section {
	p := New(strings.Split(strings.TrimSuffix(body, "\n"), "\n"))

	for i, line := range p.lines {
		p.line(i+1, strings.TrimSpace(line))
	}

	p.closeSection(len(p.lines))

	return p.result
}
