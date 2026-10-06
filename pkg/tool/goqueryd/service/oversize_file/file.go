package oversize_file

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/service/oversize_chunk"

type File struct {
	Collection string
	Path       string
	Worst      int
	Chunks     []*oversize_chunk.Chunk
}
