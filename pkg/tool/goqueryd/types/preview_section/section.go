package preview_section

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/types/preview_block"

type Section struct {
	Level     int
	Title     string
	FirstLine int
	LastLine  int
	Tokens    int
	Blocks    []*preview_block.Block
}
