package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"github.com/funtimecoding/soil/pkg/source/types/import_entry"
	"github.com/funtimecoding/soil/pkg/source/types/import_view"
	"testing"
)

func TestDecideOwnConstantStaysPlain(t *testing.T) {
	d := imports.Decide(
		import_view.New(
			"example.org/alfa/pkg/web/server",
			[]*import_entry.Entry{
				entry(
					"example.org/alfa/pkg/web/constant",
					constant.ConstantPackage,
					"foxtrot",
				),
			},
			taken(),
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
				entry(
					"example.org/alfa/pkg/web/constant",
					constant.ConstantPackage,
					constant.ConstantPackage,
				),
			},
			taken(),
		),
	)
	assert.String(t, "webConstant", d[0].Want)
	assert.True(t, d[0].Kind == constant.DecisionRename)
}

func TestDecideIgnoresAnImportThatIsLeaving(t *testing.T) {
	arriving := entry(
		"example.org/alfa/pkg/tool/goalfa/constant",
		constant.ConstantPackage,
		constant.ConstantPackage,
	)
	d := imports.Decide(
		import_view.New(
			"example.org/alfa/pkg/tool/goalfa/server",
			[]*import_entry.Entry{
				leaving(
					entry(
						"example.org/alfa/pkg/system/constant",
						constant.ConstantPackage,
						constant.ConstantPackage,
					),
				),
				arriving,
			},
			taken(),
		),
	)
	assert.Count(t, 1, d)
	assert.True(t, d[0].Entry == arriving)
	assert.String(t, "constant", d[0].Want)
	assert.True(t, d[0].Kind == constant.DecisionKeep)
}

func TestDecideRefusesANameALocalHolds(t *testing.T) {
	d := imports.Decide(
		import_view.New(
			"example.org/alfa/pkg/web/server",
			[]*import_entry.Entry{
				entry(
					"example.org/alfa/pkg/web/constant",
					constant.ConstantPackage,
					"foxtrot",
				),
			},
			taken(constant.ConstantPackage),
		),
	)
	assert.True(t, d[0].Kind == constant.DecisionRefused)
	assert.String(t, "local collision with \"constant\"", d[0].Reason)
}
