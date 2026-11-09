package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestRelationStampRunsOnServiceClock(t *testing.T) {
	s := base.NewWith(t, windowNow)
	c := s.MemoryClient(t)
	target, e := c.PostMemory(
		t.Context(),
		client.PostMemoryJSONRequestBody{
			Name:        "clocked target",
			Content:     "clocked target content",
			Description: "clocked target memory",
			Type:        new("feedback"),
		},
	)
	assert.FatalOnError(t, e)
	assert.FatalOnError(
		t,
		c.PostRelation(
			t.Context(),
			client.PostRelationJSONRequestBody{
				SourceIdentifier: clockedMemory(t, s),
				TargetIdentifier: target.Identifier,
			},
		),
	)
	relations, f := s.Store().ListRelations()
	assert.FatalOnError(t, f)
	assert.Count(t, 1, relations)
	assert.String(t, "2026-08-16T20:00:00Z", relations[0].CreatedAt)
}

func TestMemoryCreateStampRunsOnServiceClock(t *testing.T) {
	s := base.NewWith(t, windowNow)
	m, e := s.Store().GetMemory(clockedMemory(t, s))
	assert.FatalOnError(t, e)
	assert.String(t, "2026-08-16T20:00:00Z", m.CreatedAt)
}

func TestMemoryUpdateStampRunsOnServiceClock(t *testing.T) {
	s := base.NewWith(t, windowNow)
	identifier := clockedMemory(t, s)
	_, e := s.MemoryClient(t).PutMemory(
		t.Context(),
		identifier,
		client.PutMemoryJSONRequestBody{Content: new("changed content")},
	)
	assert.FatalOnError(t, e)
	m, f := s.Store().GetMemory(identifier)
	assert.FatalOnError(t, f)
	assert.String(t, "2026-08-16T20:00:00Z", m.UpdatedAt)
}

func TestMemoryForgetStampRunsOnServiceClock(t *testing.T) {
	s := base.NewWith(t, windowNow)
	identifier := clockedMemory(t, s)
	_, e := s.MemoryClient(t).DeleteMemory(
		t.Context(),
		identifier,
		&client.DeleteMemoryParams{},
	)
	assert.FatalOnError(t, e)
	m, f := s.Store().GetMemory(identifier)
	assert.FatalOnError(t, f)
	assert.False(t, m.IsActive)
	assert.String(t, "2026-08-16T20:00:00Z", m.UpdatedAt)
}

func TestImpressionArrivalStampRunsOnServiceClock(t *testing.T) {
	s := base.NewWith(t, windowNow)
	c, e := client.NewClientWithResponses(
		s.Address(webConstant.RootPath),
		client.WithHTTPClient(s.Client()),
		client.WithRequestEditorFn(
			web.BearerEditor(generativeConstant.ModelContextTestToken),
		),
	)
	assert.FatalOnError(t, e)
	_, f := c.PostImpressionsWithResponse(
		t.Context(),
		client.PostImpressionsJSONRequestBody{Content: "arrived impression"},
	)
	assert.FatalOnError(t, f)
	impressions, g := s.Store().RecentImpressions("")
	assert.FatalOnError(t, g)
	assert.Count(t, 1, impressions)
	assert.String(t, "2026-08-16T20:00:00Z", impressions[0].CreatedAt)
}
