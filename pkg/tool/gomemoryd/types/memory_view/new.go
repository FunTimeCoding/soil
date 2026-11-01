package memory_view

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func New(m *record.Memory) *View {
	return &View{Memory: m}
}
