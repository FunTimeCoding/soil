package sweep

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"path/filepath"
)

func DeleteSource(
	directories []string,
	sessionIdentifier string,
) []string {
	var result []string
	target := join.Empty(sessionIdentifier, constant.NotationLogExtension)

	for _, d := range directories {
		path := filepath.Join(d, target)

		if f := os.Remove(path); f == nil {
			result = append(result, path)
		}

		directory := filepath.Join(d, sessionIdentifier)

		if _, f := os.Stat(directory); f != nil {
			continue
		}

		if f := os.RemoveAll(directory); f == nil {
			result = append(result, directory)
		}
	}

	return result
}
