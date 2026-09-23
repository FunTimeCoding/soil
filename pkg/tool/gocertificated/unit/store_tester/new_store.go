package store_tester

import (
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
)

func NewStore() *store.Store {
	return store.New(lite.NewMemory())
}
