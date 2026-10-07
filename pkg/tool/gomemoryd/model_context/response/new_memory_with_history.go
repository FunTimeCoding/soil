package response

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func NewMemoryWithHistory(
	memory record.Memory,
	related []record.Related,
	history []record.Version,
) *MemoryWithHistory {
	return &MemoryWithHistory{
		Memory:  memory,
		Related: related,
		History: history,
	}
}
