package layer

import (
	"bytes"
	"os"
)

func (l *Layer) stale() []string {
	var result []string

	for current := l; current != nil; current = current.parent {
		for path, expected := range current.expected {
			disk, e := os.ReadFile(path)

			if e != nil || !bytes.Equal(disk, expected) {
				result = append(result, path)
			}
		}
	}

	return result
}
