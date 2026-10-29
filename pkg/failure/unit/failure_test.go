package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

type alfa interface {
	Alfa() error
}

func TestAFailureReachesOnlyTheMethodItNames(t *testing.T) {
	f := NewFixture()
	f.Fail(f.Bravo, errors.New("bravo refused"))
	assert.Nil(t, f.Alfa())
	assert.String(t, "bravo refused", f.Bravo().Error())
}

func TestAFailureSetOnOneMockReachesNoOther(t *testing.T) {
	f := NewFixture()
	g := NewFixture()
	f.Fail(f.Alfa, errors.New("alfa refused"))
	assert.Nil(t, g.Alfa())
}

func TestAMethodThroughAnInterfaceIsRefused(t *testing.T) {
	f := NewFixture()
	var a alfa = f
	assert.Panic(t, func() { f.Fail(a.Alfa, errors.New("never")) })
}
