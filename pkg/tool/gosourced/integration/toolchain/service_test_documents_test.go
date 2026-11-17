package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"strings"
	"testing"
)

func TestRenameTestRetargetsBareCitation(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("rename-follow/src"))
	system_tester.Write(
		t,
		o.Directory,
		"pkg/shop/unit/shop_test.go",
		"package unit\n\nimport \"testing\"\n\nfunc TestShelfHoldsStock(t *testing.T) {}\n",
	)
	system_tester.Write(
		t,
		o.Directory,
		"doc/note.md",
		"Pinned by `TestShelfHoldsStock`.\n",
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/shop/unit",
		"TestShelfHoldsStock",
		"TestShelfKeepsStock",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"Pinned by `TestShelfKeepsStock`.\n",
		o.Read("doc/note.md"),
	)
}

func TestRenameRepeatedTestLeavesCitationForHand(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("rename-follow/src"))

	for _, p := range []string{
		"pkg/shop/unit/shop_test.go",
		"pkg/stock/unit/stock_test.go",
	} {
		system_tester.Write(
			t,
			o.Directory,
			p,
			"package unit\n\nimport \"testing\"\n\nfunc TestGuard(t *testing.T) {}\n",
		)
	}

	system_tester.Write(
		t,
		o.Directory,
		"doc/note.md",
		"Pinned by `TestGuard`.\n",
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/shop/unit",
		"TestGuard",
		"TestShopGuard",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	assert.String(t, "Pinned by `TestGuard`.\n", o.Read("doc/note.md"))
	reported := false

	for _, c := range r.Entries {
		if strings.Contains(c.Text, "declared in another package too") {
			reported = true
		}
	}

	assert.True(t, reported)
}
