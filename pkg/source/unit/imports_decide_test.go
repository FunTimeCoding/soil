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

func TestDecideOwnConstantStaysPlain(t *testing.T) {
	d := imports.Decide(
		import_view.New(
			"example.org/alfa/pkg/web/server",
			[]*import_entry.Entry{
				import_tester.Entry(
					"example.org/alfa/pkg/web/constant",
					constant.ConstantPackage,
					"foxtrot",
				),
			},
			import_tester.Among(),
		),
	)
	assert.Count(t, 1, d)
	assert.String(t, "constant", d[0].Want)
	assert.True(t, d[0].Kind == constant.DecisionRename)
}

func TestDecideForeignConstantCarriesItsDomain(t *testing.T) {
	d := imports.Decide(
		import_view.New(
			"example.org/alfa/pkg/tool/goalfa",
			[]*import_entry.Entry{
				import_tester.Entry(
					"example.org/alfa/pkg/web/constant",
					constant.ConstantPackage,
					constant.ConstantPackage,
				),
			},
			import_tester.Among(),
		),
	)
	assert.String(t, "webConstant", d[0].Want)
	assert.True(t, d[0].Kind == constant.DecisionRename)
}

func TestDecideRefusesNameLaterConstantKeeps(t *testing.T) {
	d := imports.Decide(
		import_view.New(
			"example.org/alfa/pkg/tool/goalfa",
			[]*import_entry.Entry{
				import_tester.Entry(
					"example.org/alfa/pkg/web/constant",
					constant.ConstantPackage,
					constant.ConstantPackage,
				),
				import_entry.New(
					"example.org/bravo/pkg/web/constant",
					constant.ConstantPackage,
					"webConstant",
					"example.org/bravo",
					false,
				),
			},
			import_tester.Among(),
		),
	)
	assert.Count(t, 2, d)
	assert.True(t, d[0].Kind == constant.DecisionRefused)
	assert.String(t, "webConstant already given in this file", d[0].Reason)
	assert.True(t, d[1].Kind == constant.DecisionKeep)
}

func TestDecideIgnoresImportThatIsLeaving(t *testing.T) {
	arriving := import_tester.Entry(
		"example.org/alfa/pkg/tool/goalfa/constant",
		constant.ConstantPackage,
		constant.ConstantPackage,
	)
	d := imports.Decide(
		import_view.New(
			"example.org/alfa/pkg/tool/goalfa/server",
			[]*import_entry.Entry{
				import_tester.Leaving(
					import_tester.Entry(
						"example.org/alfa/pkg/system/constant",
						constant.ConstantPackage,
						constant.ConstantPackage,
					),
				),
				arriving,
			},
			import_tester.Among(),
		),
	)
	assert.Count(t, 1, d)
	assert.True(t, d[0].Entry == arriving)
	assert.String(t, "constant", d[0].Want)
	assert.True(t, d[0].Kind == constant.DecisionKeep)
}

func TestDecideLeavesLosersWhenKeeperIsRefused(t *testing.T) {
	d := imports.Decide(
		import_view.New(
			"example.org/alfa/pkg/tool/goalfa/face/service",
			[]*import_entry.Entry{
				import_tester.Entry(
					"example.org/alfa/pkg/tool/goalfa/definition/object",
					"object",
					"definition",
				),
				import_tester.Entry(
					"example.org/alfa/pkg/tool/goalfa/store/object",
					"object",
					"repository",
				),
			},
			import_tester.Among("object"),
		),
	)
	assert.Count(t, 2, d)
	assert.True(t, d[0].Kind == constant.DecisionRefused)
	assert.String(t, "local collision with \"object\"", d[0].Reason)
	assert.True(t, d[1].Kind == constant.DecisionKeep)
	assert.String(t, "repository", d[1].Want)
}

func TestDecideRefusesNameLocalHolds(t *testing.T) {
	d := imports.Decide(
		import_view.New(
			"example.org/alfa/pkg/web/server",
			[]*import_entry.Entry{
				import_tester.Entry(
					"example.org/alfa/pkg/web/constant",
					constant.ConstantPackage,
					"foxtrot",
				),
			},
			import_tester.Among(constant.ConstantPackage),
		),
	)
	assert.True(t, d[0].Kind == constant.DecisionRefused)
	assert.String(t, "local collision with \"constant\"", d[0].Reason)
}
