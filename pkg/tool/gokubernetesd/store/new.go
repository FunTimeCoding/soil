package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store/mute_rule"
)

func New(path string) *Store {
	d := lite.New(path)
	errors.PanicOnError(d.AutoMigrate(mute_rule.Stub()))

	return &Store{database: d}
}
