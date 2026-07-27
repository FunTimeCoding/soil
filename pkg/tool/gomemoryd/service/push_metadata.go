package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
)

func pushMetadata(m *store.Memory) map[string]string {
	result := map[string]string{}

	for key, value := range m.Metadata {
		result[key] = value
	}

	result[constant.MemoryIdentifier] = fmt.Sprintf("%d", m.Identifier)
	result[constant.Type] = m.Type
	result[constant.MemoryName] = m.Name
	result[constant.Description] = m.Description

	if m.Scope != "" {
		result[constant.Scope] = m.Scope
	}

	return result
}
