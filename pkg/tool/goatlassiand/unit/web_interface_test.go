package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/service_tester"
	"testing"
)

func TestPlateSectionAffirmsWhenClean(t *testing.T) {
	assert.StringContains(
		t,
		"All clean.",
		base.New(t).Body(t, constant.PlatePath),
	)
}

func TestEmptySectionsHideTheirHeadings(t *testing.T) {
	body := base.New(t).Body(t, constant.PlatePath)
	assert.StringNotContains(t, "Watched Issues", body)
	assert.StringNotContains(t, "Favourites", body)
	assert.StringNotContains(t, "Watched Pages", body)
}

func TestHiddenSectionsKeepTheirLiveTargets(t *testing.T) {
	body := base.New(t).Body(t, constant.PlatePath)
	assert.StringContains(t, `sse-swap="plate"`, body)
	assert.StringContains(t, `sse-swap="watched_issues"`, body)
	assert.StringContains(t, `sse-swap="favorites"`, body)
	assert.StringContains(t, `sse-swap="watched_pages"`, body)
}

func TestNewestSectionAbsentWithoutConfiguredProject(t *testing.T) {
	body := base.New(t).Body(t, constant.PlatePath)
	assert.StringNotContains(t, "Newest Issues", body)
	assert.StringNotContains(t, `sse-swap="newest"`, body)
}

func TestNewestSectionKeepsLiveTargetWithConfiguredProject(t *testing.T) {
	body := base.New(
		t,
		service_tester.WithProject([]string{"ABC"}),
	).Body(t, constant.PlatePath)
	assert.StringContains(t, `sse-swap="newest"`, body)
	assert.StringNotContains(t, "Newest Issues", body)
}

func TestClockFillsPlatePage(t *testing.T) {
	s := base.New(t)
	s.Refresh(service_tester.OpenIssue("ALFA-1"))
	body := s.Body(t, constant.PlatePath)
	assert.StringContains(t, "ALFA-1", body)
	assert.StringNotContains(t, "All clean.", body)
}
