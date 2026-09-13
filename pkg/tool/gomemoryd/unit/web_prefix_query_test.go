package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web"
	"testing"
)

func TestPrefixQueryAppendsToLastToken(t *testing.T) {
	assert.String(t, "scr*", web.PrefixQuery("scr"))
	assert.String(t, "memory scr*", web.PrefixQuery("memory scr"))
}

func TestPrefixQueryLeavesExistingStar(t *testing.T) {
	assert.String(t, "scr*", web.PrefixQuery("scr*"))
}

func TestPrefixQueryLeavesEmpty(t *testing.T) {
	assert.String(t, "", web.PrefixQuery(""))
	assert.String(t, "   ", web.PrefixQuery("   "))
}

func TestPrefixQueryLeavesOperatorAndPunctuation(t *testing.T) {
	assert.String(t, `foo "bar"`, web.PrefixQuery(`foo "bar"`))
	assert.String(t, "foo (", web.PrefixQuery("foo ("))
	assert.String(t, "foo:", web.PrefixQuery("foo:"))
}

func TestPrefixQueryKeepsHyphenAndDigits(t *testing.T) {
	assert.String(t, "utf-8*", web.PrefixQuery("utf-8"))
	assert.String(t, "sha256*", web.PrefixQuery("sha256"))
}
