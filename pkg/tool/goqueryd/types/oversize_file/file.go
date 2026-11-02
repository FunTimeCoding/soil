package oversize_file

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/types/oversize_chunk"

type File struct {
	Collection string
	Path       string
	Worst      int
	Chunks     []*oversize_chunk.Chunk
}
