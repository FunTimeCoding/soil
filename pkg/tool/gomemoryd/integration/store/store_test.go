package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"testing"
)

func TestCreateAndGetMemory(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Content = "Retry failed requests with exponential backoff."
	o.Description = "retry with backoff; each caller gets its own policy"
	o.Type = "feedback"
	o.Tags = []string{"always", "http-client"}
	identifier := s.CreateMemory(o)
	m := s.GetMemory(identifier)
	assert.String(
		t,
		"Retry failed requests with exponential backoff.",
		m.Content,
	)
	assert.Count(t, 2, m.Tags)
}

func TestForgetMemory(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Content = "to forget"
	o.Description = "desc"
	o.Type = "feedback"
	identifier := s.CreateMemory(o)
	s.ForgetMemory(identifier, "test")
	assert.Count(t, 0, s.ListMemories("", "", "", true))
	assert.Count(t, 1, s.ListMemories("", "", "", false))
}

func TestCreateMemoryWithProvenance(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Name = "Retry"
	o.Content = "Broken draws are retried."
	o.Description = ""
	o.Type = "reference"
	o.Scope = "alfa"
	o.Metadata = map[string]string{"kind": "mechanism", "guard": "true"}
	o.ProvenanceFile = "canon/Example.yaml"
	o.ProvenanceAnchor = "Retry"
	o.ProvenanceHash = "abc123"
	o.Ordinal = 4
	identifier := s.CreateMemory(o)
	m := s.GetMemory(identifier)
	assert.String(t, "canon/Example.yaml", m.ProvenanceFile)
	assert.String(t, "Retry", m.ProvenanceAnchor)
	assert.String(t, "abc123", m.ProvenanceHash)
	assert.Integer(t, 4, m.Ordinal)
	assert.String(t, "mechanism", m.Metadata["kind"])
	assert.String(t, "true", m.Metadata["guard"])
}

func TestListDocumentSourced(t *testing.T) {
	s := store_tester.New(t)
	plain := save_option.New()
	plain.Name = "hand-tended"
	plain.Content = "no provenance"
	plain.Description = "plain"
	plain.Type = "feedback"
	s.CreateMemory(plain)
	o := save_option.New()
	o.Name = "Example"
	o.Content = "parent"
	o.Description = "file parent"
	o.Type = "reference"
	o.Scope = "alfa"
	o.ProvenanceFile = "canon/Example.yaml"
	o.ProvenanceHash = "parent-hash"
	parent := s.CreateMemory(o)
	child := save_option.New()
	child.Name = "Shard"
	child.Content = "shard text"
	child.Type = "reference"
	child.Scope = "alfa"
	child.ParentIdentifier = &parent
	child.ProvenanceFile = "canon/Example.yaml"
	child.ProvenanceAnchor = "Shard"
	child.ProvenanceHash = "shard-hash"
	child.Ordinal = 1
	s.CreateMemory(child)
	sourced, e := s.Store.ListDocumentSourced("alfa")
	assert.FatalOnError(t, e)
	assert.Count(t, 2, sourced)
	assert.String(t, "Example", sourced[0].Name)
	assert.String(t, "parent-hash", sourced[0].ProvenanceHash)
	assert.String(t, "Shard", sourced[1].Name)
	assert.Integer(t, 1, sourced[1].Ordinal)
	empty, f := s.Store.ListDocumentSourced("")
	assert.FatalOnError(t, f)
	assert.Count(t, 0, empty)
}

func TestUpdateMemoryReplacesMetadataAndOrdinal(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Name = "Shard"
	o.Content = "first text"
	o.Type = "reference"
	o.Scope = "alfa"
	o.Metadata = map[string]string{"kind": "mechanism"}
	o.ProvenanceFile = "canon/Example.yaml"
	o.ProvenanceAnchor = "Shard"
	o.ProvenanceHash = "hash-one"
	o.Ordinal = 1
	identifier := s.CreateMemory(o)
	update := save_option.New()
	update.Name = "Shard"
	update.Content = "second text"
	update.Metadata = map[string]string{"kind": "growth"}
	update.ProvenanceHash = "hash-two"
	update.Ordinal = 3
	s.UpdateMemory(identifier, update)
	m := s.GetMemory(identifier)
	assert.String(t, "second text", m.Content)
	assert.String(t, "growth", m.Metadata["kind"])
	assert.String(t, "hash-two", m.ProvenanceHash)
	assert.Integer(t, 3, m.Ordinal)
	assert.String(t, "canon/Example.yaml", m.ProvenanceFile)
	assert.String(t, "Shard", m.ProvenanceAnchor)
}

func TestChildrenOrderedByOrdinal(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Name = "parent"
	o.Content = "parent content"
	o.Description = "parent"
	o.Type = "reference"
	parent := s.CreateMemory(o)
	second := save_option.New()
	second.Name = "second shard"
	second.Content = "second"
	second.Description = "second"
	second.Type = "reference"
	second.ParentIdentifier = &parent
	second.Ordinal = 2
	s.CreateMemory(second)
	first := save_option.New()
	first.Name = "first shard"
	first.Content = "first"
	first.Description = "first"
	first.Type = "reference"
	first.ParentIdentifier = &parent
	first.Ordinal = 1
	s.CreateMemory(first)
	children, e := s.Store.ListChildren(parent)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, children)
	assert.String(t, "first shard", children[0].Name)
	assert.String(t, "second shard", children[1].Name)
}

func TestListRelated(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Name = "error handling"
	o.Content = "error handling content"
	o.Description = "error handling description"
	o.Type = "feedback"
	id1 := s.CreateMemory(o)
	p := save_option.New()
	p.Name = "deployment"
	p.Content = "deployment content"
	p.Description = "deployment description"
	p.Type = "feedback"
	p.Tags = []string{"deploy", "no-index"}
	id2 := s.CreateMemory(p)
	q := save_option.New()
	q.Name = "telemetry"
	q.Content = "telemetry content"
	q.Description = "telemetry description"
	q.Type = "project"
	id3 := s.CreateMemory(q)
	s.CreateRelation(id1, id2)
	s.CreateRelation(id3, id1)
	related := s.ListRelated(id1)
	assert.Count(t, 2, related)
	assert.Integer(t, id2, related[0].Identifier)
	assert.String(t, "deployment", related[0].Name)
	assert.String(t, "deployment description", related[0].Description)
	assert.Count(t, 2, related[0].Tags)
	assert.Integer(t, id3, related[1].Identifier)
	assert.String(t, "telemetry", related[1].Name)
	s.ForgetMemory(id2, "test")
	assert.Count(t, 1, s.ListRelated(id1))
}

func TestScopeSeparatesListing(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Name = "default memory"
	o.Content = "Lives in the default scope."
	o.Description = "default"
	o.Type = "feedback"
	s.CreateMemory(o)
	p := save_option.New()
	p.Name = "scoped memory"
	p.Content = "Lives in a named scope."
	p.Description = "scoped"
	p.Type = "reference"
	p.Scope = "alfa"
	s.CreateMemory(p)
	defaults := s.ListMemories("", "", "", true)
	assert.Count(t, 1, defaults)
	assert.String(t, "default memory", defaults[0].Name)
	scoped := s.ListMemories("", "", "alfa", true)
	assert.Count(t, 1, scoped)
	assert.String(t, "scoped memory", scoped[0].Name)
	assert.String(t, "alfa", scoped[0].Scope)
	all := s.ListMemories("", "", constant.AllScope, true)
	assert.Count(t, 2, all)
}

func TestScopeSeparatesSearch(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Name = "default note"
	o.Content = "The turbine spins in the default scope."
	o.Description = "default turbine"
	o.Type = "feedback"
	s.CreateMemory(o)
	p := save_option.New()
	p.Name = "scoped note"
	p.Content = "The turbine spins in a named scope."
	p.Description = "scoped turbine"
	p.Type = "reference"
	p.Scope = "alfa"
	s.CreateMemory(p)
	defaults := s.SearchMemories("turbine", 10, "", "", "")
	assert.Count(t, 1, defaults)
	assert.String(t, "default note", defaults[0].Name)
	scoped := s.SearchMemories("turbine", 10, "", "", "alfa")
	assert.Count(t, 1, scoped)
	assert.String(t, "scoped note", scoped[0].Name)
	all := s.SearchMemories("turbine", 10, "", "", constant.AllScope)
	assert.Count(t, 2, all)
}

func TestSearchMemories(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Content = "Always validate input at service boundaries"
	o.Description = "validate at boundaries; trust internal callers"
	o.Type = "feedback"
	o.Tags = []string{"always"}
	s.CreateMemory(o)
	p := save_option.New()
	p.Content = "Build regression tests to verify fix correctness"
	p.Description = "regression tests over manual verification"
	p.Type = "feedback"
	s.CreateMemory(p)
	results := s.SearchMemories("validate input", 10, "", "", "")
	assert.Greater(t, 0, len(results))
	assert.String(
		t,
		"validate at boundaries; trust internal callers",
		results[0].Description,
	)
}

func TestListTags(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Content = "a"
	o.Description = "desc a"
	o.Type = "feedback"
	o.Tags = []string{"always", "deployment"}
	s.CreateMemory(o)
	p := save_option.New()
	p.Content = "b"
	p.Description = "desc b"
	p.Type = "feedback"
	p.Tags = []string{"always"}
	s.CreateMemory(p)
	tags := s.ListTags()
	assert.Count(t, 2, tags)
	assert.String(t, "always", tags[0].Tag)
	assert.Integer(t, 2, tags[0].Count)
}

func TestUpdateCreatesVersion(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Content = "original"
	o.Description = "desc"
	o.Type = "feedback"
	identifier := s.CreateMemory(o)
	p := save_option.New()
	p.Content = "updated content"
	p.Description = "new desc"
	p.Tags = []string{"always"}
	s.UpdateMemory(identifier, p)
	history := s.GetMemoryHistory(identifier)
	assert.Count(t, 2, history)
	assert.String(t, "created", history[0].ChangeType)
	assert.String(t, "updated", history[1].ChangeType)
	m := s.GetMemory(identifier)
	assert.String(t, "updated content", m.Content)
	assert.Count(t, 1, m.Tags)
	assert.String(t, "always", m.Tags[0])
}
