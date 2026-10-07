package response

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
)

func NewSlimMemoryWithHistory(
	slimMemory convert.SlimMemory,
	related []record.Related,
	history []record.Version,
) *SlimMemoryWithHistory {
	return &SlimMemoryWithHistory{
		SlimMemory: slimMemory,
		Related:    related,
		History:    history,
	}
}
