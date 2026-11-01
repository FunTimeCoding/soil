package memory_view

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

type View struct {
	Memory  *record.Memory
	History []record.Version
	Related []record.Related
}
