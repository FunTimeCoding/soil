package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestLocatorTargetSplitsTheQuery(t *testing.T) {
	assert.String(
		t,
		"https://example.org/detail?id=1",
		locator.New(constant.Example).Target("/detail?id=1").String(),
	)
	assert.String(
		t,
		"https://example.org/entries?system=a&system=b&user=bob",
		locator.New(constant.Example).Target(
			"/entries?user=bob&system=a&system=b",
		).String(),
	)
}

func TestLocatorTargetWithoutQueryIsAPath(t *testing.T) {
	assert.String(
		t,
		"https://example.org/health",
		locator.New(constant.Example).Target(constant.HealthPath).String(),
	)
}

func TestLocatorTargetKeepsSetValues(t *testing.T) {
	assert.String(
		t,
		"https://example.org/detail?id=1&page=2",
		locator.New(constant.Example).Set("page", "2").Target(
			"/detail?id=1",
		).String(),
	)
}

func TestLocatorTargetRefusesAMalformedQuery(t *testing.T) {
	assert.Panic(
		t,
		func() { locator.New(constant.Example).Target("/detail?id=%zz") },
	)
}
