package match

import "github.com/funtimecoding/soil/pkg/tool/goreplace/block"

type Match struct {
	Block  *block.Block
	Offset int
	Line   int
}
