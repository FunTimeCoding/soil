package lite

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/lite/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewMemory() *gorm.DB {
	m, e := gorm.Open(
		sqlite.Open(join.Empty(constant.Memory, constant.MemoryParameters)),
		&gorm.Config{},
	)
	errors.PanicOnError(e)
	inner, f := m.DB()
	errors.PanicOnError(f)
	// A pooled second connection would open its own empty in-memory database
	inner.SetMaxOpenConns(1)

	return m
}
