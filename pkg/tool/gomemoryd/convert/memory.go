package convert

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func Memory(m *store.Memory) *SlimMemory {
	return &SlimMemory{
		Identifier:       m.Identifier,
		Name:             m.Name,
		Content:          m.Content,
		Description:      m.Description,
		Tags:             m.Tags,
		Metadata:         m.Metadata,
		ParentIdentifier: m.ParentIdentifier,
		Ordinal:          m.Ordinal,
	}
}
