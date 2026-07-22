package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
	"time"
)

type Fixture struct {
	Value string
}

func TestAny(t *testing.T) {
	assert.Any(t, &Fixture{Value: "a"}, &Fixture{Value: "a"})
}

func TestBoolean(t *testing.T) {
	assert.Boolean(t, true, true)
}

func TestBytes(t *testing.T) {
	assert.Bytes(t, []byte("123"), []byte("123"))
}

func TestDuration(t *testing.T) {
	assert.Duration(t, 5*time.Second, 5*time.Second)
}

func TestFalse(t *testing.T) {
	assert.False(t, false)
}

func TestFloat(t *testing.T) {
	assert.Float(t, 1.1, 1.1)
	assert.Float(t, 1.15, 1.15)
}

func TestInteger(t *testing.T) {
	assert.Integer(t, 1, 1)
}

func TestInteger32(t *testing.T) {
	assert.Integer32(t, 1, 1)
}

func TestInteger64(t *testing.T) {
	assert.Integer64(t, 1, 1)
}

func TestNil(t *testing.T) {
	assert.Nil(t, nil)
}

func TestNotEmpty(t *testing.T) {
	assert.NotEmpty(t, []string{"one"})
	assert.NotEmpty(t, map[string]int{"a": 1})
}

func TestNotNil(t *testing.T) {
	assert.NotNil(t, 1)
}

func TestString(t *testing.T) {
	assert.String(t, "a", "a")
}

func TestTime(t *testing.T) {
	now := time.Now()
	assert.Time(t, now, now)
}

func TestTrue(t *testing.T) {
	assert.True(t, true)
}

func TestUnsigned(t *testing.T) {
	assert.Unsigned(t, 1, 1)
}

func TestUnsigned32(t *testing.T) {
	assert.Unsigned32(t, 1, 1)
}

func TestNewDay(t *testing.T) {
	assert.Time(t, time.Date(1970, 1, 2, 0, 0, 0, 0, time.UTC), assert.NewDay(2))
}

func TestNewMinute(t *testing.T) {
	assert.Time(
		t,
		time.Date(2000, 1, 1, 0, 5, 0, 0, time.UTC),
		assert.NewMinute(5),
	)
}

func TestStub(t *testing.T) {
	before := assert.StubCount
	assert.Stub(t)
	assert.Integer(t, before+1, assert.StubCount)
}
