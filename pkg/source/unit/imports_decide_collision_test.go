package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"github.com/funtimecoding/soil/pkg/source/types/import_entry"
	"github.com/funtimecoding/soil/pkg/source/types/import_view"
	"github.com/funtimecoding/soil/pkg/source/unit/import_tester"
	"testing"
)

func TestDecideGivesEveryCollidingMockItsOwner(t *testing.T) {
	d := imports.Decide(
		import_view.New(
			"example.org/alfa/pkg/tool/goalfa/unit",
			[]*import_entry.Entry{
				import_tester.Entry(
					"example.org/alfa/pkg/bravo/mock_client",
					constant.MockPackage,
					constant.MockPackage,
				),
				import_tester.Entry(
					"example.org/alfa/pkg/tool/goalfa/client/mock_client",
					constant.MockPackage,
					"alfaMock",
				),
			},
			import_tester.Among(),
		),
	)
	assert.Count(t, 2, d)
	assert.String(t, "bravoMock", d[0].Want)
	assert.True(t, d[0].Kind == constant.DecisionRename)
	assert.String(t, "goalfaMock", d[1].Want)
	assert.True(t, d[1].Kind == constant.DecisionRename)
}

func TestDecideKeepsThirdPartyLosersAlias(t *testing.T) {
	v := import_view.New(
		"example.org/alfa/pkg/alert/store",
		[]*import_entry.Entry{
			import_tester.Entry("example.org/alfa/pkg/alert", "alert", "alert"),
			import_entry.New(
				"example.net/upstream/client/alert",
				"alert",
				"raw",
				"example.net/upstream",
				false,
			),
		},
		import_tester.Among(),
	)
	v.Ours = func(modulePath string) bool {
		return modulePath == "example.org/alfa"
	}
	d := imports.Decide(v)
	assert.Count(t, 2, d)
	assert.True(t, d[1].Kind == constant.DecisionKeep)
	assert.String(t, "raw", d[1].Want)
}

func TestDecideNeverGivesTwoGroupsOneName(t *testing.T) {
	d := imports.Decide(
		import_view.New(
			"example.org/alfa/pkg/tool/goalfa",
			[]*import_entry.Entry{
				import_tester.Entry(
					"example.org/alfa/pkg/tool/goalfa/client",
					"client",
					"client",
				),
				import_tester.Entry(
					"example.org/alfa/pkg/tool/gobravo/client",
					"client",
					"bravoClient",
				),
				import_tester.Entry(
					"example.org/alfa/pkg/tool/goalfa/face",
					"face",
					"face",
				),
				import_tester.Entry(
					"example.org/alfa/pkg/tool/gobravo/face",
					"face",
					"bravo",
				),
			},
			import_tester.Among(),
		),
	)
	assert.Count(t, 4, d)
	assert.String(t, "gobravo", d[1].Want)
	assert.String(t, "gobravoFace", d[3].Want)
}

func TestDecideKeepsAnotherGroupsPlainNameFree(t *testing.T) {
	d := imports.Decide(
		import_view.New(
			"example.org/alfa/pkg/tool/goalfa",
			[]*import_entry.Entry{
				import_tester.Entry(
					"example.org/alfa/pkg/tool/goalfa/client",
					"client",
					"alfaClient",
				),
				import_tester.Entry(
					"example.org/alfa/pkg/tool/gobravo/client",
					"client",
					"bravoClient",
				),
				import_tester.Entry(
					"example.org/alfa/pkg/tool/goalfa/face",
					"face",
					"face",
				),
				import_tester.Entry(
					"example.org/alfa/pkg/charlie/client/face",
					"face",
					"charlie",
				),
			},
			import_tester.Among(),
		),
	)
	assert.Count(t, 4, d)
	assert.String(t, "clientFace", d[3].Want)
}

func TestDecideRenamesFirstPartyLoser(t *testing.T) {
	d := imports.Decide(
		import_view.New(
			"example.org/alfa/pkg/alert/store",
			[]*import_entry.Entry{
				import_tester.Entry(
					"example.org/alfa/pkg/alert",
					"alert",
					"alert",
				),
				import_tester.Entry(
					"example.org/alfa/pkg/upstream/alert",
					"alert",
					"raw",
				),
			},
			import_tester.Among(),
		),
	)
	assert.Count(t, 2, d)
	assert.True(t, d[1].Kind == constant.DecisionRename)
	assert.String(t, "upstream", d[1].Want)
}
