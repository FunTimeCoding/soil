package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"testing"
)

func TestServiceCreateMemory(t *testing.T) {
	o := service_tester.New(t)
	p := save_option.New()
	p.Name = "retry policy"
	p.Content = "Retry failed requests with exponential backoff."
	p.Description = "retry with backoff"
	p.Source = "test"
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)
	assert.String(t, "retry policy", m.Name)
	assert.String(t, "feedback", m.Type)
	assert.Count(t, 1, o.Indexer.Pushed)
	assert.String(t, "memory/1", o.Indexer.Pushed[0].Name)
}

func TestServiceCreateMemoryWithExplicitType(t *testing.T) {
	o := service_tester.New(t)
	p := save_option.New()
	p.Name = "deploy target"
	p.Content = "Production deployments use blue-green strategy."
	p.Description = "blue-green deploy pattern"
	p.Type = "user"
	p.Source = "test"
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)
	assert.String(t, "user", m.Type)
}

func TestServiceForgetMemory(t *testing.T) {
	o := service_tester.New(t)
	p := save_option.New()
	p.Name = "to forget"
	p.Content = constant.FixtureContent
	p.Description = "desc"
	p.Type = "feedback"
	p.Source = "test"
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, o.Service.ForgetMemory(m.Identifier, "test"))
	assert.Count(t, 1, o.Indexer.Deleted)
	assert.String(t, "memory/1", o.Indexer.Deleted[0].Path)
	assert.String(t, "memories", o.Indexer.Deleted[0].Collection)
	active, e := o.Service.ListMemories("", "", "", true)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, active)
}

func TestHiddenTagLeavesIndex(t *testing.T) {
	o := service_tester.New(t)
	o.Service.WithHiddenTag("private")
	p := save_option.New()
	p.Name = "delta"
	p.Content = "quiet content"
	p.Description = "quiet description"
	p.Type = "project"
	p.Source = "test"
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, o.Indexer.Pushed)
	assert.FatalOnError(t, o.Service.AddTags(m.Identifier, []string{"private"}))
	assert.Count(t, 1, o.Indexer.Pushed)
	assert.Count(t, 1, o.Indexer.Deleted)
	q := save_option.New()
	q.Name = "delta"
	q.Content = "updated quiet content"
	q.Description = "quiet description"
	q.Source = "test"
	_, _, e = o.Service.UpdateMemory(m.Identifier, q)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, o.Indexer.Pushed)
	assert.Count(t, 2, o.Indexer.Deleted)
	assert.FatalOnError(
		t,
		o.Service.RemoveTags(m.Identifier, []string{"private"}),
	)
	assert.Count(t, 2, o.Indexer.Pushed)
	hidden, e := o.Service.HiddenIdentifiers()
	assert.FatalOnError(t, e)
	assert.Count(t, 0, hidden)
}

func TestRetiredMemoryStaysOutOfTheIndexOnTagEdit(t *testing.T) {
	o := service_tester.New(t)
	p := save_option.New()
	p.Name = "alfa"
	p.Content = "first"
	p.Description = "a test"
	p.Source = "test"
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, o.Service.ForgetMemory(m.Identifier, "test"))
	pushed := len(o.Indexer.Pushed)
	assert.FatalOnError(
		t,
		o.Service.ReplaceTags(m.Identifier, []string{"build"}),
	)
	assert.Integer(t, pushed, len(o.Indexer.Pushed))
	assert.String(
		t,
		"memory/1",
		o.Indexer.Deleted[len(o.Indexer.Deleted)-1].Path,
	)
}

func TestActiveMemoryIsIndexedOnTagEdit(t *testing.T) {
	o := service_tester.New(t)
	p := save_option.New()
	p.Name = "bravo"
	p.Content = "second"
	p.Description = "a test"
	p.Source = "test"
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)
	pushed := len(o.Indexer.Pushed)
	assert.FatalOnError(
		t,
		o.Service.ReplaceTags(m.Identifier, []string{"build"}),
	)
	assert.Integer(t, pushed+1, len(o.Indexer.Pushed))
}

func TestServiceCreateMemoryRejectsReservedScope(t *testing.T) {
	o := service_tester.New(t)

	for _, reserved := range []string{
		constant.AllScope,
		constant.DefaultScope,
	} {
		_, e := o.Service.CreateMemory(scopedOption("reserved", reserved))
		assert.True(t, validation.Is(e))
	}
}

func TestServiceProfileRejectsReservedScope(t *testing.T) {
	o := service_tester.New(t)
	_, _, e := o.Service.Profile("", constant.AllScope, false)
	assert.True(t, validation.Is(e))
}

func TestServiceCreateRoutesCollectionByScope(t *testing.T) {
	o := service_tester.New(t)
	_, e := o.Service.CreateMemory(scopedOption("default entry", ""))
	assert.FatalOnError(t, e)
	assert.String(t, "memories", o.Indexer.Pushed[0].Collection)
	_, f := o.Service.CreateMemory(scopedOption("scoped entry", "alfa"))
	assert.FatalOnError(t, f)
	assert.Count(t, 2, o.Indexer.Pushed)
	assert.String(t, "alfa", o.Indexer.Pushed[1].Collection)
	scope := o.Indexer.Pushed[1].Metadata[constant.Scope]
	assert.Count(t, 1, scope)
	assert.String(t, "alfa", scope[0])
}

func TestServiceUpdatePreservesMetadataAndOrdinal(t *testing.T) {
	o := service_tester.New(t)
	p := scopedOption("scoped entry", "alfa")
	p.Metadata = map[string]string{"kind": "mechanism"}
	p.Ordinal = 2
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)
	q := save_option.New()
	q.Name = "scoped entry"
	q.Content = "updated"
	q.Description = "scoped entry"
	q.Source = "test"
	updated, _, f := o.Service.UpdateMemory(m.Identifier, q)
	assert.FatalOnError(t, f)
	assert.String(t, "mechanism", updated.Metadata["kind"])
	assert.Integer(t, 2, updated.Ordinal)
	assert.String(t, "alfa", o.Indexer.Pushed[1].Collection)
}

func TestServiceCreateMemoryWithBase(t *testing.T) {
	o := service_tester.New(t)
	p := scopedOption("error capture", "")
	p.Base = new("../soil/doc/ai/spec/error-handling")
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)
	assert.String(
		t,
		"../soil/doc/ai/spec/error-handling",
		m.Metadata[constant.BaseKey],
	)
}

func TestServiceUpdateKeepsBaseWhenOmitted(t *testing.T) {
	o := service_tester.New(t)
	m := createBased(t, o)
	updated, _, e := o.Service.UpdateMemory(m.Identifier, updateOption())
	assert.FatalOnError(t, e)
	assert.String(t, "doc", updated.Metadata[constant.BaseKey])
	assert.String(t, "mechanism", updated.Metadata["kind"])
}

func TestServiceUpdateClearsBaseWhenEmpty(t *testing.T) {
	o := service_tester.New(t)
	m := createBased(t, o)
	q := updateOption()
	q.Base = new("")
	updated, _, e := o.Service.UpdateMemory(m.Identifier, q)
	assert.FatalOnError(t, e)
	_, found := updated.Metadata[constant.BaseKey]
	assert.Boolean(t, false, found)
	assert.String(t, "mechanism", updated.Metadata["kind"])
}

func TestServiceUpdateReplacesBase(t *testing.T) {
	o := service_tester.New(t)
	m := createBased(t, o)
	q := updateOption()
	q.Base = new("pkg")
	updated, _, e := o.Service.UpdateMemory(m.Identifier, q)
	assert.FatalOnError(t, e)
	assert.String(t, "pkg", updated.Metadata[constant.BaseKey])
	assert.String(t, "mechanism", updated.Metadata["kind"])
}

func TestServiceProfileScoped(t *testing.T) {
	o := service_tester.New(t)
	_, e := o.Service.CreateMemory(scopedOption("default entry", ""))
	assert.FatalOnError(t, e)
	_, f := o.Service.CreateMemory(scopedOption("scoped entry", "alfa"))
	assert.FatalOnError(t, f)
	parent, g := o.Service.GetMemory(2)
	assert.FatalOnError(t, g)
	second := scopedOption("second shard", "alfa")
	second.ParentIdentifier = &parent.Identifier
	second.Ordinal = 2
	_, h := o.Service.CreateMemory(second)
	assert.FatalOnError(t, h)
	first := scopedOption("first shard", "alfa")
	first.ParentIdentifier = &parent.Identifier
	first.Ordinal = 1
	_, i := o.Service.CreateMemory(first)
	assert.FatalOnError(t, i)
	result, _, g := o.Service.Profile("", "alfa", false)
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

func TestServiceUpdateMemoryPreservesTags(t *testing.T) {
	o := service_tester.New(t)
	p := save_option.New()
	p.Name = "pace"
	p.Content = "original content"
	p.Description = "original desc"
	p.Type = "feedback"
	p.Source = "test"
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)
	assert.FatalOnError(
		t,
		o.Service.AddTags(m.Identifier, []string{"always", "go-conventions"}),
	)
	q := save_option.New()
	q.Name = "pace"
	q.Content = "updated content"
	q.Description = "updated desc"
	q.Source = "test"
	updated, _, e := o.Service.UpdateMemory(m.Identifier, q)
	assert.FatalOnError(t, e)
	assert.String(t, "updated content", updated.Content)
	assert.Count(t, 2, updated.Tags)
	assert.Count(t, 3, o.Indexer.Pushed)
}

func TestServiceUpdateMemoryNonexistentFails(t *testing.T) {
	o := service_tester.New(t)
	p := save_option.New()
	p.Name = constant.FixtureName
	p.Content = constant.FixtureContent
	p.Description = "desc"
	p.Source = "test"
	_, _, e := o.Service.UpdateMemory(999, p)
	assert.Error(t, e)
}

func TestServiceUpdateWritesGivenEmptyContent(t *testing.T) {
	o := service_tester.New(t)
	m := createNamed(t, o, "pace", "", "Original body.")
	q := renameOption(m, "pace")
	q.Content = ""
	updated, _, e := o.Service.UpdateMemory(m.Identifier, q)
	assert.FatalOnError(t, e)
	assert.String(t, "", updated.Content)
	assert.String(t, "pace", updated.Description)
}

func TestServiceRenameRewritesCitations(t *testing.T) {
	o := service_tester.New(t)
	m := createNamed(t, o, "pace", "", "Original body.")
	citing := createNamed(
		t,
		o,
		"citing",
		"",
		"See `memory://default/pace` first.",
	)
	q := renameOption(m, "tempo")
	_, rewritten, e := o.Service.UpdateMemory(m.Identifier, q)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, rewritten)
	assert.String(t, "citing", rewritten[0].Name)
	reloaded, f := o.Service.GetMemory(citing.Identifier)
	assert.FatalOnError(t, f)
	assert.String(t, "See `memory://default/tempo` first.", reloaded.Content)
}

func TestServiceRenameLeavesSimilarNames(t *testing.T) {
	o := service_tester.New(t)
	m := createNamed(t, o, "pace", "", "Original body.")
	citing := createNamed(
		t,
		o,
		"citing",
		"",
		"See `memory://default/pace of work` first.",
	)
	q := renameOption(m, "tempo")
	_, rewritten, e := o.Service.UpdateMemory(m.Identifier, q)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, rewritten)
	reloaded, f := o.Service.GetMemory(citing.Identifier)
	assert.FatalOnError(t, f)
	assert.String(
		t,
		"See `memory://default/pace of work` first.",
		reloaded.Content,
	)
}

func TestServiceRenameFollowsScope(t *testing.T) {
	o := service_tester.New(t)
	m := createNamed(t, o, "pace", "alfa", "Original body.")
	citing := createNamed(
		t,
		o,
		"citing",
		"",
		"See `memory://alfa/pace` and `memory://default/pace`.",
	)
	q := renameOption(m, "tempo")
	_, _, e := o.Service.UpdateMemory(m.Identifier, q)
	assert.FatalOnError(t, e)
	reloaded, f := o.Service.GetMemory(citing.Identifier)
	assert.FatalOnError(t, f)
	assert.String(
		t,
		"See `memory://alfa/tempo` and `memory://default/pace`.",
		reloaded.Content,
	)
}

func TestServiceRenameOntoExistingNameIsRefused(t *testing.T) {
	o := service_tester.New(t)
	m := createNamed(t, o, "pace", "", "Original body.")
	createNamed(t, o, "tempo", "", "Another body.")
	q := renameOption(m, "tempo")
	_, _, e := o.Service.UpdateMemory(m.Identifier, q)
	assert.True(t, validation.Is(e))
	reloaded, f := o.Service.GetMemory(m.Identifier)
	assert.FatalOnError(t, f)
	assert.String(t, "pace", reloaded.Name)
}

func TestServiceRenameVersionsTheRewrittenMemory(t *testing.T) {
	o := service_tester.New(t)
	m := createNamed(t, o, "pace", "", "Original body.")
	citing := createNamed(t, o, "citing", "", "See `memory://default/pace`.")
	q := renameOption(m, "tempo")
	_, _, e := o.Service.UpdateMemory(m.Identifier, q)
	assert.FatalOnError(t, e)
	history, f := o.Service.GetMemoryHistory(citing.Identifier)
	assert.FatalOnError(t, f)
	last := history[len(history)-1]
	assert.String(t, "rewritten", last.ChangeType)
	assert.String(t, "test", last.Source)
	assert.String(t, "See `memory://default/tempo`.", last.Content)
}
