package store_tester

import (
	"github.com/funtimecoding/soil/pkg/time/clock_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"testing"
)

type Tester struct {
	Store *store.Store
	t     *testing.T
	clock *clock_tester.Clock
}
