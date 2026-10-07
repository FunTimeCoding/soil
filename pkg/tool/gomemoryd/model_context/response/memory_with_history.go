package response

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

type MemoryWithHistory struct {
	record.Memory
	Related []record.Related `json:"related,omitempty"`
	History []record.Version `json:"history,omitempty"`
}
