package unit

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type Fixture struct {
	Value string
}

type ExportedFixture struct {
	Value  string
	hidden string
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
	assert.Integer(t, 1, 1)
}

func TestInteger64(t *testing.T) {
	assert.Integer(t, 1, 1)
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
	assert.Integer(t, 1, 1)
}

func TestUnsigned32(t *testing.T) {
	assert.Integer(t, 1, 1)
}

func TestNewDay(t *testing.T) {
	assert.Time(
		t,
		time.Date(1970, 1, 2, 0, 0, 0, 0, time.UTC),
		assert.NewDay(2),
	)
}

func TestNewMinute(t *testing.T) {
	assert.Time(
		t,
		time.Date(2000, 1, 1, 0, 5, 0, 0, time.UTC),
		assert.NewMinute(5),
	)
}

func TestContains(t *testing.T) {
	assert.Contains(t, []string{}, []string{"a", "b"})
	assert.Contains(t, []string{"a"}, []string{"a", "b"})
	assert.Contains(t, []string{"a", "b"}, []string{"a", "b"})
}

func TestCount(t *testing.T) {
	assert.Count(t, 0, []string{})
	assert.Count(t, 1, []string{"1"})
	assert.Count(t, 2, []string{"1", "2"})
}

func TestIntegers(t *testing.T) {
	assert.Integers(t, []int{0, 1}, []int{0, 1})
}

func TestMapHasKey(t *testing.T) {
	m := map[string]any{"height": 180}
	assert.MapHasKey(t, m, "height")
}

func TestMapNotHasKey(t *testing.T) {
	m := map[string]any{"height": 180}
	assert.MapNotHasKey(t, m, "weight")
}

func TestMapValue(t *testing.T) {
	m := map[string]any{"height": 180, "name": "Alfa"}
	assert.MapValue(t, 180, m, "height")
	assert.MapValue(t, "Alfa", m, "name")
}

func TestPrefix(t *testing.T) {
	assert.Prefix(t, "", "ab")
	assert.Prefix(t, "a", "ab")
	assert.Prefix(t, "ab", "ab")
}

func TestStringContains(t *testing.T) {
	assert.StringContains(t, "friend", "hello friend")
}

func TestStringNotContains(t *testing.T) {
	assert.StringNotContains(t, "enemy", "hello friend")
}

func TestStrings(t *testing.T) {
	assert.Strings(t, []string{"a", "b"}, []string{"a", "b"})
}

func TestSuffix(t *testing.T) {
	assert.Suffix(t, "", "ab")
	assert.Suffix(t, "b", "ab")
	assert.Suffix(t, "ab", "ab")
}

func TestError(t *testing.T) {
	assert.Error(t, fmt.Errorf("something went wrong"))
}

func TestErrorIs(t *testing.T) {
	sentinel := errors.New("not found")
	assert.ErrorIs(t, fmt.Errorf("thing not found: %w", sentinel), sentinel)
}

func TestFatalOnError(t *testing.T) {
	assert.FatalOnError(t, nil)
}

func TestExported(t *testing.T) {
	assert.Exported(
		t,
		&ExportedFixture{Value: "a"},
		&ExportedFixture{Value: "a"},
	)
}

func TestExportedIgnoresPrivateFields(t *testing.T) {
	assert.Exported(
		t,
		&ExportedFixture{Value: "a", hidden: "b"},
		&ExportedFixture{Value: "a", hidden: "c"},
	)
}

func TestHTTPStatus(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				q *http.Request,
			) {
				w.WriteHeader(http.StatusTeapot)
			},
		),
	)
	defer server.Close()
	assert.HTTPStatus(t, server.URL, http.StatusTeapot)
}

func TestDeviate(t *testing.T) {
	assert.Deviate(t, 1, 1.1, 0.1)
	assert.Deviate(t, -1, -1.1, 0.1)
}

func TestGreater(t *testing.T) {
	assert.Greater(t, 0, 1)
}

func TestGreaterRejectsLess(t *testing.T) {
	inner := &testing.T{}
	assert.Greater(inner, 1, 0)
	assert.True(t, inner.Failed())
}

func TestGreaterRejectsEqual(t *testing.T) {
	inner := &testing.T{}
	assert.Greater(inner, 1, 1)
	assert.True(t, inner.Failed())
}

func TestGreaterEqual(t *testing.T) {
	assert.GreaterEqual(t, 0, 1)
	assert.GreaterEqual(t, 1, 1)
}

func TestGreaterEqualRejectsLess(t *testing.T) {
	inner := &testing.T{}
	assert.GreaterEqual(inner, 1, 0)
	assert.True(t, inner.Failed())
}

func TestLess(t *testing.T) {
	assert.Less(t, 1, 0)
}

func TestLessRejectsMore(t *testing.T) {
	inner := &testing.T{}
	assert.Less(inner, 0, 1)
	assert.True(t, inner.Failed())
}

func TestLessRejectsEqual(t *testing.T) {
	inner := &testing.T{}
	assert.Less(inner, 1, 1)
	assert.True(t, inner.Failed())
}

func TestLessEqual(t *testing.T) {
	assert.LessEqual(t, 1, 0)
	assert.LessEqual(t, 1, 1)
}

func TestLessEqualRejectsMore(t *testing.T) {
	inner := &testing.T{}
	assert.LessEqual(inner, 0, 1)
	assert.True(t, inner.Failed())
}

func TestRound(t *testing.T) {
	assert.Round(t, 1.1, 1.14, 1)
	assert.Round(t, 1.15, 1.154, 2)
}

func TestFixture(t *testing.T) {
	assert.Suffix(t, "soil/fixture/example.txt", fixture.Path("example.txt"))
}

func TestFileExists(t *testing.T) {
	assert.FileExists(t, "./assert_test.go")
}
