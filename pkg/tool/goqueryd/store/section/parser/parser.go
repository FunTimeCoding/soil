package parser

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/section"

type Parser struct {
	lines      []string
	result     []*section.Section
	current    *section.Section
	blockStart int
	blockKind  string
	inFence    bool
	inFront    bool
}
