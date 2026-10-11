package checksum

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func Path(workDirectory string) string {
	return join.Absolute(workDirectory, constant.ChecksumFile)
}
