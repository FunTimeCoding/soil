package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service/oversize_file"
)

func convertOversizeFile(f *oversize_file.File) server.OversizeFile {
	chunks := make([]server.OversizeChunk, len(f.Chunks))

	for i, c := range f.Chunks {
		chunks[i] = server.OversizeChunk{
			Index:     c.Index,
			FirstLine: c.FirstLine,
			LastLine:  c.LastLine,
			Bytes:     c.Bytes,
			Tokens:    c.Tokens,
		}
	}

	return server.OversizeFile{
		Collection: f.Collection,
		Path:       f.Path,
		Worst:      f.Worst,
		Chunks:     chunks,
	}
}
