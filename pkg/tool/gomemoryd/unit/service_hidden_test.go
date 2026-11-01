package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/service_tester"
	"testing"
)

func TestDashboardViewCountsOnlyVisibleMemories(t *testing.T) {
	o := service_tester.New(t)
	o.Service.WithHiddenTag("private")
	alfa := createTagged(t, o, "alfa", "shared")
	bravo := createTagged(t, o, "bravo", "shared", "private")
	charlie := createTagged(t, o, "charlie")
	assert.FatalOnError(
		t,
		o.Service.CreateRelation(alfa.Identifier, bravo.Identifier, ""),
	)
	assert.FatalOnError(
		t,
		o.Service.CreateRelation(alfa.Identifier, charlie.Identifier, ""),
	)
	v, e := o.Service.DashboardView()
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, v.Active)
	assert.Integer(t, 2, v.Types["feedback"])
	assert.Count(t, 1, v.Tags)
	assert.String(t, "shared", v.Tags[0].Tag)
	assert.Integer(t, 1, v.Tags[0].Count)
	assert.Integer(t, 1, v.Relations)
	assert.Integer(t, 1, v.Untyped)
	assert.Count(t, 2, v.Versions)

	for _, version := range v.Versions {
		assert.True(t, version.Name != "bravo")
	}
}

func TestMemoryListLeavesHiddenMemoriesOut(t *testing.T) {
	o := service_tester.New(t)
	o.Service.WithHiddenTag("private")
	createTagged(t, o, "alfa")
	createTagged(t, o, "bravo", "private")
	l, e := o.Service.MemoryList("", "", "", 1)
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, l.Total)
	assert.Count(t, 1, l.Memories)
	assert.String(t, "alfa", l.Memories[0].Name)
}

func TestHiddenMemoryViewReadsAsNotFound(t *testing.T) {
	o := service_tester.New(t)
	o.Service.WithHiddenTag("private")
	bravo := createTagged(t, o, "bravo", "private")
	_, e := o.Service.MemoryView(bravo.Identifier)
	assert.True(t, not_found.Is(e))
}

func TestMemoryViewLeavesHiddenRelatedOut(t *testing.T) {
	o := service_tester.New(t)
	o.Service.WithHiddenTag("private")
	alfa := createTagged(t, o, "alfa")
	bravo := createTagged(t, o, "bravo", "private")
	charlie := createTagged(t, o, "charlie")
	assert.FatalOnError(
		t,
		o.Service.CreateRelation(alfa.Identifier, bravo.Identifier, ""),
	)
	assert.FatalOnError(
		t,
		o.Service.CreateRelation(alfa.Identifier, charlie.Identifier, ""),
	)
	v, e := o.Service.MemoryView(alfa.Identifier)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v.Related)
	assert.String(t, "charlie", v.Related[0].Name)
}

func TestVisibleRelationsLeaveHiddenEndpointsOut(t *testing.T) {
	o := service_tester.New(t)
	o.Service.WithHiddenTag("private")
	alfa := createTagged(t, o, "alfa")
	bravo := createTagged(t, o, "bravo", "private")
	charlie := createTagged(t, o, "charlie")
	assert.FatalOnError(
		t,
		o.Service.CreateRelation(bravo.Identifier, alfa.Identifier, ""),
	)
	assert.FatalOnError(
		t,
		o.Service.CreateRelation(alfa.Identifier, charlie.Identifier, ""),
	)
	relations, e := o.Service.VisibleRelations()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, relations)
	assert.String(t, "charlie", relations[0].TargetName)
}
