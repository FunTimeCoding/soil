package board

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"gopkg.in/yaml.v3"
)

func Load(path string) *Board {
	result := &Board{}
	errors.PanicOnError(yaml.Unmarshal(system.ReadBytesUnsafe(path), result))

	if result.Tail.Columns == 0 {
		result.Tail.Columns = constant.DefaultTailColumns
	}

	result.validate()

	return result
}
