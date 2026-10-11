package module

import "github.com/funtimecoding/soil/pkg/go_mod/constant"

func (m *Module) formatUntagged() string {
	if m.Untagged {
		return constant.Untagged
	}

	return ""
}
