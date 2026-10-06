package model_context

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

type memoryWithHistory struct {
	record.Memory
	Related []record.Related `json:"related,omitempty"`
	History []record.Version `json:"history,omitempty"`
}
