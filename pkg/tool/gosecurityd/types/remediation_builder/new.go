package remediation_builder

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_item"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
)

func New() *Builder {
	return &Builder{
		items:       map[string]*remediation_item.Item{},
		hosts:       map[[2]string]bool{},
		found:       map[[2]string]bool{},
		worst:       map[[2]string]*system_result.Finding{},
		unimportant: map[string]bool{},
	}
}
