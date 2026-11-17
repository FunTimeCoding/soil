package oversize

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/types/oversize_file"

type Report struct {
	Model           string
	Window          int
	TokenLimit      int
	EmbedTokenLimit int
	Files           []*oversize_file.File
}
