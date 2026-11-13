package module_graph

import (
	"io/fs"
	"os"
)

func diskDirectory(directory string) ([]fs.FileInfo, error) {
	entries, e := os.ReadDir(directory)

	if e != nil {
		return nil, e
	}

	result := make([]fs.FileInfo, 0, len(entries))

	for _, d := range entries {
		i, f := d.Info()

		if f != nil {
			return nil, f
		}

		result = append(result, i)
	}

	return result, nil
}
