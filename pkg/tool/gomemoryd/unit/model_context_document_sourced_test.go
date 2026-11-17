package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"testing"
	"time"
)

func TestDocumentSourcedMemoryRejectsUpdate(t *testing.T) {
	s := base.New(t)
	c := s.Tools()
	identifier := s.DocumentSourcedMemory()
	result := c.MustCallToolError(
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
	s := base.New(t)
	c := s.Tools()
	identifier := s.DocumentSourcedMemory()
	result := c.MustCallToolError(
		constant.ForgetMemory,
		map[string]any{constant.MemoryIdentifier: identifier},
	)
	assert.StringContains(t, "document-sourced", result)
	m, e := s.Store().GetMemory(identifier)
	assert.FatalOnError(t, e)
	assert.True(t, m.IsActive)
}

func TestDocumentSourcedMemoryRejectsTag(t *testing.T) {
	s := base.New(t)
	c := s.Tools()
	identifier := s.DocumentSourcedMemory()
	result := c.MustCallToolError(
		constant.TagMemory,
		map[string]any{
			constant.MemoryIdentifier: identifier,
			constant.Add:              "extra",
		},
	)
	assert.StringContains(t, "document-sourced", result)
}

func TestGetMemoryGroup(t *testing.T) {
	s := base.New(t)
	c := s.Tools()
	o := save_option.New()
	o.Name = "Cluster"
	o.Content = "parent"
	o.Description = "cluster parent"
	o.Type = "reference"
	parent, e := s.Store().CreateMemory(o, time.Now())
	assert.FatalOnError(t, e)
	child := save_option.New()
	child.Name = "member"
	child.Content = "child content"
	child.Description = "member"
	child.Type = "reference"
	child.ParentIdentifier = &parent
	child.Ordinal = 1
	_, f := s.Store().CreateMemory(child, time.Now())
	assert.FatalOnError(t, f)
	result := c.MustCallTool(
		constant.GetMemoryGroup,
		map[string]any{constant.MemoryIdentifier: parent},
	)
	assert.StringContains(t, "Cluster", result)
	assert.StringContains(t, "child content", result)
}
