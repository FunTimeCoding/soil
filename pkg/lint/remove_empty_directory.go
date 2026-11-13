package lint

import "os"

func removeEmptyDirectory(path string) bool {
	return os.Remove(path) == nil
}
