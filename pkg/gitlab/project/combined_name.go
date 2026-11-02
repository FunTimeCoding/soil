package project

import (
	"fmt"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
)

func (p *Project) CombinedName() string {
	return fmt.Sprintf("%s%s%s", p.Namespace, stringsConstant.Slash, p.Name)
}
