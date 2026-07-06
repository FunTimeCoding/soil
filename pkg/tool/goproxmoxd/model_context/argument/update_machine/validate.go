package update_machine

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"strings"
)

func (m *Machine) Validate() error {
	setNames := make(map[string]bool)

	if m.Name != "" {
		setNames["name"] = true
	}

	if m.Tags != "" {
		setNames["tags"] = true
	}

	if m.OnBoot != nil {
		setNames["onboot"] = true
	}

	if m.Cores > 0 {
		setNames["cores"] = true
	}

	if m.Memory > 0 {
		setNames["memory"] = true
	}

	if m.Description != "" {
		setNames["description"] = true
	}

	if m.Delete != "" {
		for _, field := range strings.Split(m.Delete, ",") {
			if setNames[strings.TrimSpace(field)] {
				return constant.ErrorSetDeleteConflict
			}
		}
	}

	if len(setNames) == 0 && m.Delete == "" {
		return constant.ErrorNoChanges
	}

	return nil
}
