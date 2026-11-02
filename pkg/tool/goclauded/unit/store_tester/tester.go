package store_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"gorm.io/gorm"
	"testing"
	"time"
)

type Tester struct {
	t      *testing.T
	mapper *gorm.DB
	Store  *store.Store
	now    *time.Time
}
