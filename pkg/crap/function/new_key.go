package function

import (
	"fmt"
	"path"
	"path/filepath"
)

func NewKey(
	packagePath string,
	file string,
	line int,
) string {
	return fmt.Sprintf(
		"%s:%d",
		path.Join(packagePath, filepath.Base(file)),
		line,
	)
}
