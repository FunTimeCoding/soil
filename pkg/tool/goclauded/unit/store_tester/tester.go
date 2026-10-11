package store_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"testing"
	"time"
)

type Tester struct {
	t     *testing.T
	Store *store.Store
	now   *time.Time
}
