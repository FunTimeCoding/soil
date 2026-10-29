package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/unit/service_tester"
	"testing"
)

func TestServiceClosesDatabase(t *testing.T) {
	s := service_tester.Service(t)
	_, e := s.Count()
	assert.FatalOnError(t, e)
	s.Close()
	_, f := s.Count()
	assert.Error(t, f)
}
