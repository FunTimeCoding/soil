package store_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store"
	"testing"
)

type Tester struct {
	t     *testing.T
	Store *store.Store
}
