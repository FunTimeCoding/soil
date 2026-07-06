package project

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/separator"
)

func (p *Project) CombinedName() string {
	return fmt.Sprintf("%s%s%s", p.Namespace, separator.Slash, p.Name)
}
