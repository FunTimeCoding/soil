package store_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/option"
	"path/filepath"
	"testing"
	"time"
)

func NewOption(
	t *testing.T,
	clock func() time.Time,
	n face.EventNotifier,
) *option.Sprout {
	t.Helper()
	result := option.New()
	result.LitePath = filepath.Join(t.TempDir(), "gosproutd.db")
	result.Clock = clock
	result.Notifier = func() face.EventNotifier { return n }

	return result
}
