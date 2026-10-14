package helper

import "sync"

type Resource struct{}

func Open() *Resource {
	return &Resource{}
}

func (r *Resource) Close() {}

type Holder struct {
	Resource *Resource
}

func (h *Holder) Close() {
	h.Resource.Close()
}

type Leaky struct {
	Resource *Resource
}

func (l *Leaky) Close() {}

func (r *Resource) Name() string {
	return ""
}

var Shared = sync.OnceValue(Open)
