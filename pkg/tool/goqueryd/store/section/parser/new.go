package parser

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/section"
	"strings"
)

func New(lines []string) *Parser {
	return &Parser{
		lines:   lines,
		current: section.NewSection(1),
		inFront: len(lines) > 0 &&
			strings.TrimSpace(lines[0]) == constant.FrontMatterFence,
	}
}
