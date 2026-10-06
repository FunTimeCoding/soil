package oversize

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/service/oversize_file"

type Report struct {
	Model     string
	Window    int
	Allowance int
	Files     []*oversize_file.File
}
