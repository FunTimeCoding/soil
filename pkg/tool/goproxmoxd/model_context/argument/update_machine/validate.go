package update_machine

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"strings"
)

func (m *Machine) Validate() error {
	setNames := make(map[string]bool)

	if m.Name != "" {
		setNames[constant.NameOption] = true
	}

	if m.Tags != "" {
		setNames[constant.TagsOption] = true
	}

	if m.OnBoot != nil {
		setNames[constant.OnBootOption] = true
	}

	if m.Cores > 0 {
		setNames[constant.CoresOption] = true
	}

	if m.Memory > 0 {
		setNames[constant.MemoryOption] = true
	}

	if m.Description != "" {
		setNames[constant.DescriptionOption] = true
	}

	if m.Delete != "" {
		for _, field := range strings.Split(m.Delete, ",") {
			if setNames[strings.TrimSpace(field)] {
				return validation.New("cannot set and delete the same field")
			}
		}
	}

	if len(setNames) == 0 && m.Delete == "" {
		return validation.New("no changes specified")
	}

	return nil
}
