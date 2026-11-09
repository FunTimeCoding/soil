package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/limit_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/service_tester"
	"testing"
)

func TestAbsentLogLimitReadsOneHundred(t *testing.T) {
	c := limit_recorder.New()
	defer c.Close()
	_, e := service_tester.New(t, c).Log(0)
	assert.FatalOnError(t, e)
	assert.Integer(t, 100, c.Limit())
}

func TestGivenLogLimitIsKept(t *testing.T) {
	c := limit_recorder.New()
	defer c.Close()
	_, e := service_tester.New(t, c).Log(7)
	assert.FatalOnError(t, e)
	assert.Integer(t, 7, c.Limit())
}
