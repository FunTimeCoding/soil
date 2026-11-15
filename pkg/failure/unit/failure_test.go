package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/failure/unit/failure_tester"
	"testing"
)

type alfa interface {
	Alfa() error
}

func TestFailureReachesOnlyMethodItNames(t *testing.T) {
	f := failure_tester.New()
	f.Fail(f.Bravo, errors.New("bravo refused"))
	assert.Nil(t, f.Alfa())
	assert.String(t, "bravo refused", f.Bravo().Error())
}

func TestFailureSetOnOneMockReachesNoOther(t *testing.T) {
	f := failure_tester.New()
	g := failure_tester.New()
	f.Fail(f.Alfa, errors.New("alfa refused"))
	assert.Nil(t, g.Alfa())
}

func TestMethodThroughInterfaceIsRefused(t *testing.T) {
	f := failure_tester.New()
	var a alfa = f
	assert.Panic(t, func() { f.Fail(a.Alfa, errors.New("never")) })
}
