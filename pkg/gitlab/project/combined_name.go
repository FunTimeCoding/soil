package project

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/constant"
)

func (p *Project) CombinedName() string {
	return fmt.Sprintf("%s%s%s", p.Namespace, constant.Slash, p.Name)
}
