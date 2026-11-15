package store_tester

import (
	"github.com/funtimecoding/soil/pkg/time/clock_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"gorm.io/gorm"
	"testing"
)

type Tester struct {
	t      *testing.T
	mapper *gorm.DB
	Store  *store.Store
	clock  *clock_tester.Clock
}
