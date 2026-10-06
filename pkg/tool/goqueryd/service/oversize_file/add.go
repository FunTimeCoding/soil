package oversize_file

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/service/oversize_chunk"

func (f *File) Add(c *oversize_chunk.Chunk) {
	f.Chunks = append(f.Chunks, c)
	f.Worst = max(f.Worst, c.Tokens)
}
