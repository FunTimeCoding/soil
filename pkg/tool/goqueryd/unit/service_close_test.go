package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"testing"
)

func TestTheServiceClosesWhatItBuiltOnce(t *testing.T) {
	m := newEmbedder(t)
	a := newReranker(t)
	v := service.New(
		service_tester.NewOption(t, m, a),
		logger.New(t.Context()),
		memory.New(),
	)
	v.Close()
	v.Close()
	assert.Integer(t, 1, m.Closed())
	assert.Integer(t, 1, a.Closed)
	_, e := v.Status()
	assert.Error(t, e)
}
