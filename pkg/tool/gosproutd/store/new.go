package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/seed"
)

func New(path string) *Store {
	d := lite.New(path)
	errors.PanicOnError(d.AutoMigrate(seed.Stub()))

	return &Store{mapper: d}
}
