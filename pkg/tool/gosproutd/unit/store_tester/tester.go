package store_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"testing"
	"time"
)

type Tester struct {
	Store *store.Store
	t     *testing.T
	now   *time.Time
}
