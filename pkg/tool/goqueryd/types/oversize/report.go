package oversize

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/types/oversize_file"

type Report struct {
	Model          string
	Window         int
	Allowance      int
	EmbedAllowance int
	Files          []*oversize_file.File
}
