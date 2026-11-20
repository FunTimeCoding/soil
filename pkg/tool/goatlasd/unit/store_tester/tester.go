package store_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store"
	"testing"
)

type Tester struct {
	Store *store.Store
	t     *testing.T
}
