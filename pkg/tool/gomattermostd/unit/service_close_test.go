package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"testing"
)

func TestServiceClosesWhatItBuilt(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	_, k := worker_tester.NewSink(t)
	s := service_tester.Service(t, m, k, nil)
	_, e := s.Subscriptions("kilo")
	assert.FatalOnError(t, e)
	s.Close()
	assert.Integer(t, 1, m.Closed)
	_, f := s.Subscriptions("kilo")
	assert.Error(t, f)
	s.Close()
	assert.Integer(t, 1, m.Closed)
}
