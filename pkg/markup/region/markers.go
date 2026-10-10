package region

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/markup/constant"
)

func markers(name string) (string, string) {
	return fmt.Sprintf(constant.RegionStartFormat, name),
		fmt.Sprintf(constant.RegionEndFormat, name)
}
