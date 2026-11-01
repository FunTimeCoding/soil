package memory_list

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

type List struct {
	Memories []record.MemorySummary
	Total    int
	Page     int
	Limit    int
	Scopes   []string
	Type     string
	Tag      string
	Scope    string
}
