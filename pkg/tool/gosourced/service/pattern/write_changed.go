package pattern

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"os"
)

func writeChanged(
	decorations *decoration.Set,
	resolver *resolve.Names,
	out *sink.Sink,
) (map[string]bool, error) {
	result := map[string]bool{}

	for filename, file := range decorations.Files {
		content, e := decoration.Restored(
			resolver,
			decorations.PackagePaths[file],
			map[string]string{},
			file,
		)

		if e != nil {
			return nil, e
		}

		if before, f := os.ReadFile(filename); f == nil &&
			bytes.Equal(before, content) {
			continue
		}

		result[filename] = true
		out.Write(filename, content)
	}

	return result, nil
}
