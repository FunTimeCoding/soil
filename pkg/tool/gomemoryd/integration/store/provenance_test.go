package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"testing"
)

func TestCreateMemoryWithProvenance(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Name = "Retry"
	o.Content = "Broken draws are retried."
	o.Description = ""
	o.Type = "reference"
	o.Scope = "alpha"
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
	o.Scope = "alpha"
	o.ProvenanceFile = "canon/Example.yaml"
	o.ProvenanceHash = "parent-hash"
	parent := s.CreateMemory(o)
	child := save_option.New()
	child.Name = "Shard"
	child.Content = "shard text"
	child.Type = "reference"
	child.Scope = "alpha"
	child.ParentIdentifier = &parent
	child.ProvenanceFile = "canon/Example.yaml"
	child.ProvenanceAnchor = "Shard"
	child.ProvenanceHash = "shard-hash"
	child.Ordinal = 1
	s.CreateMemory(child)
	sourced, e := s.Store.ListDocumentSourced("alpha")
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
	o.Scope = "alpha"
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
