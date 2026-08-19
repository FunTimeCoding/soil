package service

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration_test/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"testing"
)

func scopedOption(
	name string,
	scope string,
) *save_option.Option {
	o := save_option.New()
	o.Name = name
	o.Content = constant.FixtureContent
	o.Description = name
	o.Source = "test"
	o.Scope = scope

	return o
}

func TestServiceCreateMemoryRejectsReservedScope(t *testing.T) {
	o := service_tester.New(t)

	for _, reserved := range []string{constant.AllScope, constant.DefaultScope} {
		_, e := o.Service.CreateMemory(scopedOption("reserved", reserved))
		assert.True(t, errors.Is(e, constant.ErrorReservedScope))
	}
}

func TestServiceProfileRejectsReservedScope(t *testing.T) {
	o := service_tester.New(t)
	_, _, e := o.Service.Profile("", constant.AllScope, false)
	assert.True(t, errors.Is(e, constant.ErrorReservedScope))
}

func TestServiceCreateRoutesCollectionByScope(t *testing.T) {
	o := service_tester.New(t)
	_, e := o.Service.CreateMemory(scopedOption("default entry", ""))
	assert.FatalOnError(t, e)
	assert.String(t, "memories", o.Indexer.Pushed[0].Collection)
	_, f := o.Service.CreateMemory(scopedOption("scoped entry", "alpha"))
	assert.FatalOnError(t, f)
	assert.Count(t, 2, o.Indexer.Pushed)
	assert.String(t, "alpha", o.Indexer.Pushed[1].Collection)
	scope := o.Indexer.Pushed[1].Metadata[constant.Scope]
	assert.Count(t, 1, scope)
	assert.String(t, "alpha", scope[0])
}

func TestServiceUpdatePreservesMetadataAndOrdinal(t *testing.T) {
	o := service_tester.New(t)
	p := scopedOption("scoped entry", "alpha")
	p.Metadata = map[string]string{"kind": "mechanism"}
	p.Ordinal = 2
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)
	q := save_option.New()
	q.Name = "scoped entry"
	q.Content = "updated"
	q.Description = "scoped entry"
	q.Source = "test"
	updated, f := o.Service.UpdateMemory(m.Identifier, q)
	assert.FatalOnError(t, f)
	assert.String(t, "mechanism", updated.Metadata["kind"])
	assert.Integer(t, 2, updated.Ordinal)
	assert.String(t, "alpha", o.Indexer.Pushed[1].Collection)
}

func TestServiceProfileScoped(t *testing.T) {
	o := service_tester.New(t)
	_, e := o.Service.CreateMemory(scopedOption("default entry", ""))
	assert.FatalOnError(t, e)
	_, f := o.Service.CreateMemory(scopedOption("scoped entry", "alpha"))
	assert.FatalOnError(t, f)
	parent, g := o.Service.GetMemory(2)
	assert.FatalOnError(t, g)
	second := scopedOption("second shard", "alpha")
	second.ParentIdentifier = &parent.Identifier
	second.Ordinal = 2
	_, h := o.Service.CreateMemory(second)
	assert.FatalOnError(t, h)
	first := scopedOption("first shard", "alpha")
	first.ParentIdentifier = &parent.Identifier
	first.Ordinal = 1
	_, i := o.Service.CreateMemory(first)
	assert.FatalOnError(t, i)
	result, _, g := o.Service.Profile("", "alpha", false)
	assert.FatalOnError(t, g)
	assert.Count(t, 1, result.Index)
	assert.String(t, "scoped entry", result.Index[0].Name)
	assert.Count(t, 2, result.Index[0].Children)
	assert.String(t, "first shard", result.Index[0].Children[0])
	assert.String(t, "second shard", result.Index[0].Children[1])
	assert.Count(t, 0, result.Completions)
	assert.Count(t, 0, result.Impressions)
	defaultResult, _, h := o.Service.Profile("", "", false)
	assert.FatalOnError(t, h)
	assert.Count(t, 1, defaultResult.Index)
	assert.String(t, "default entry", defaultResult.Index[0].Name)
}
