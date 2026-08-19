package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
)

func pushMetadata(m *store.Memory) map[string][]string {
	result := map[string][]string{}

	for key, value := range m.Metadata {
		result[key] = []string{value}
	}

	result[constant.MemoryIdentifier] = []string{
		fmt.Sprintf("%d", m.Identifier),
	}
	result[constant.Type] = []string{m.Type}
	result[constant.MemoryName] = []string{m.Name}
	result[constant.Description] = []string{m.Description}

	if m.Scope != "" {
		result[constant.Scope] = []string{m.Scope}
	}

	return result
}
