package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"testing"
)

func TestTheServiceClosesWhatItBuilt(t *testing.T) {
	s := service_tester.New(t, t.TempDir(), t.TempDir())
	assert.Count(t, 0, s.Fights())
	s.Close()
	assert.Panic(t, func() { s.Fights() })
}
