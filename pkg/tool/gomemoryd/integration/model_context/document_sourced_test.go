package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/model_context_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"testing"
)

func documentSourcedMemory(
	t *testing.T,
	s *model_context_tester.Tester,
) int64 {
	t.Helper()
	o := save_option.New()
	o.Name = "Retry"
	o.Content = "Document-sourced content."
	o.Description = ""
	o.Type = "reference"
	o.Scope = "alpha"
	o.ProvenanceFile = "canon/Example.yaml"
	o.ProvenanceAnchor = "Retry"
	identifier, e := s.Store().CreateMemory(o)
	assert.FatalOnError(t, e)

	return identifier
}

func TestDocumentSourcedMemoryRejectsUpdate(t *testing.T) {
	s := model_context_tester.New(t)
	identifier := documentSourcedMemory(t, s)
	result := s.MustCallToolError(
		constant.UpdateMemory,
		map[string]any{
			constant.MemoryIdentifier: identifier,
			constant.MemoryName:       "Retry",
			constant.Content:          "changed",
			constant.Description:      "changed",
		},
	)
	assert.StringContains(t, "document-sourced", result)
	m, e := s.Store().GetMemory(identifier)
	assert.FatalOnError(t, e)
	assert.String(t, "Document-sourced content.", m.Content)
}

func TestDocumentSourcedMemoryRejectsForget(t *testing.T) {
	s := model_context_tester.New(t)
	identifier := documentSourcedMemory(t, s)
	result := s.MustCallToolError(
		constant.ForgetMemory,
		map[string]any{constant.MemoryIdentifier: identifier},
	)
	assert.StringContains(t, "document-sourced", result)
	m, e := s.Store().GetMemory(identifier)
	assert.FatalOnError(t, e)
	assert.True(t, m.IsActive)
}

func TestDocumentSourcedMemoryRejectsTag(t *testing.T) {
	s := model_context_tester.New(t)
	identifier := documentSourcedMemory(t, s)
	result := s.MustCallToolError(
		constant.TagMemory,
		map[string]any{
			constant.MemoryIdentifier: identifier,
			constant.Add:              "extra",
		},
	)
	assert.StringContains(t, "document-sourced", result)
}

func TestGetMemoryGroup(t *testing.T) {
	s := model_context_tester.New(t)
	o := save_option.New()
	o.Name = "Cluster"
	o.Content = "parent"
	o.Description = "cluster parent"
	o.Type = "reference"
	parent, e := s.Store().CreateMemory(o)
	assert.FatalOnError(t, e)
	child := save_option.New()
	child.Name = "member"
	child.Content = "child content"
	child.Description = "member"
	child.Type = "reference"
	child.ParentIdentifier = &parent
	child.Ordinal = 1
	_, f := s.Store().CreateMemory(child)
	assert.FatalOnError(t, f)
	result := s.MustCallTool(
		constant.GetMemoryGroup,
		map[string]any{constant.MemoryIdentifier: parent},
	)
	assert.StringContains(t, "Cluster", result)
	assert.StringContains(t, "child content", result)
}
