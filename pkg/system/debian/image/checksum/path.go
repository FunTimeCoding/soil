package checksum

import "github.com/funtimecoding/soil/pkg/system/join"

func Path(workDirectory string) string {
	return join.Absolute(workDirectory, File)
}
