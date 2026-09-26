package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/semver"
	"testing"
)

func TestTrim(t *testing.T) {
	assert.String(t, "0.0.0", semver.Trim("v0.0.0"))
}

func TestPrefixAddsTheMissingPrefix(t *testing.T) {
	assert.String(t, "v0.0.0", semver.Prefix("0.0.0"))
}

func TestPrefixKeepsAnExistingPrefix(t *testing.T) {
	assert.String(t, "v0.0.0", semver.Prefix("v0.0.0"))
}

func TestSortDescendingOrdersBareVersions(t *testing.T) {
	v := []string{"0.2.79", "0.11.152", "0.2.111", "0.0.459"}
	semver.SortDescending(v)
	assert.Strings(t, []string{"0.11.152", "0.2.111", "0.2.79", "0.0.459"}, v)
}

func TestSortDescendingOrdersPrefixedVersions(t *testing.T) {
	v := []string{"v0.2.79", "v0.11.152", "v0.2.111", "v0.0.459"}
	semver.SortDescending(v)
	assert.Strings(
		t,
		[]string{"v0.11.152", "v0.2.111", "v0.2.79", "v0.0.459"},
		v,
	)
}

func TestSortDescendingComparesMinorAboveLexical(t *testing.T) {
	v := []string{"0.2.111", "0.11.152"}
	semver.SortDescending(v)
	assert.String(t, "0.11.152", v[0])
}

func TestSortDescendingOrdersPatchNumerically(t *testing.T) {
	v := []string{"0.2.9", "0.2.111", "0.2.33"}
	semver.SortDescending(v)
	assert.Strings(t, []string{"0.2.111", "0.2.33", "0.2.9"}, v)
}
