package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/unit/service_tester"
	"testing"
)

func TestKeysBeforeConnectingAnswersNotConnected(t *testing.T) {
	_, e := service_tester.New(t).Keys()
	assert.True(t, errors.Is(e, constant.ErrorNotConnected))
}

func TestServiceReadsRunsFromItsOwnDatabase(t *testing.T) {
	v, e := service_tester.New(t).Recent(10)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, v)
}

func TestServiceClosesAStartedRunner(t *testing.T) {
	s := service_tester.New(t)

	for _, w := range s.Workers() {
		w.Start()
		t.Cleanup(w.Close)
	}
}
