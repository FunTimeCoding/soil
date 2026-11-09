package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"strings"
	"testing"
)

func TestRenameTestRetargetsBareCitation(t *testing.T) {
	d := renameFollowFixture(t)
	testutil.WriteFile(
		t,
		d,
		"pkg/shop/unit/shop_test.go",
		"package unit\n\nimport \"testing\"\n\nfunc TestShelfHoldsStock(t *testing.T) {}\n",
	)
	testutil.WriteFile(
		t,
		d,
		"doc/note.md",
		"Pinned by `TestShelfHoldsStock`.\n",
	)
	r, e := service_tester.NewService(t).Symbol.Rename(
		t.Context(),
		d,
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
		service_tester.ReadFixtureFile(t, d, "doc/note.md"),
	)
}

func TestRenameRepeatedTestLeavesCitationForHand(t *testing.T) {
	d := renameFollowFixture(t)

	for _, p := range []string{
		"pkg/shop/unit/shop_test.go",
		"pkg/stock/unit/stock_test.go",
	} {
		testutil.WriteFile(
			t,
			d,
			p,
			"package unit\n\nimport \"testing\"\n\nfunc TestGuard(t *testing.T) {}\n",
		)
	}

	testutil.WriteFile(t, d, "doc/note.md", "Pinned by `TestGuard`.\n")
	r, e := service_tester.NewService(t).Symbol.Rename(
		t.Context(),
		d,
		"example/pkg/shop/unit",
		"TestGuard",
		"TestShopGuard",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	assert.String(
		t,
		"Pinned by `TestGuard`.\n",
		service_tester.ReadFixtureFile(t, d, "doc/note.md"),
	)
	reported := false

	for _, c := range r.Entries {
		if strings.Contains(c.Text, "declared in another package too") {
			reported = true
		}
	}

	assert.True(t, reported)
}
