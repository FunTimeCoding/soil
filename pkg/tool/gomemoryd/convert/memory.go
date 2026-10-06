package convert

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func Memory(m *record.Memory) *SlimMemory {
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
