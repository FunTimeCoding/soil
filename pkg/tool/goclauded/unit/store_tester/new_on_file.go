package store_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"testing"
	"time"
)

func NewOnFile(
	t *testing.T,
	path string,
) *Tester {
	t.Helper()
	now := time.Now().UTC()
	d := lite.New(logger.New(t.Context()), path)
	t.Cleanup(func() { relational.CloseMapper(d) })

	return &Tester{
		t:     t,
		Store: store.New(d, func() time.Time { return now }),
		now:   &now,
	}
}
