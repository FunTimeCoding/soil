package model_context

import "github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"

type memoryWithHistory struct {
	store.Memory
	Related []store.Related `json:"related,omitempty"`
	History []store.Version `json:"history,omitempty"`
}
