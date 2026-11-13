package layer

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"maps"
	"strings"
)

func (l *Layer) Overlay() map[string][]byte {
	result := map[string][]byte{}

	if l.parent != nil {
		maps.Copy(result, l.parent.Overlay())
	}

	for path := range l.removed {
		if !strings.HasSuffix(path, soil.GoExtension) {
			delete(result, path)

			continue
		}

		if original := l.original[path]; original != nil {
			result[path] = stub(original)
		}
	}

	maps.Copy(result, l.written)

	if len(result) == 0 {
		return nil
	}

	return result
}
