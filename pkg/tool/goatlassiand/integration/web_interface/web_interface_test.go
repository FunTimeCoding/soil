package web_interface

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/integration/web_interface_tester"
	"testing"
)

func TestPlateSectionAffirmsWhenClean(t *testing.T) {
	o := web_interface_tester.New(t)
	o.AssertContains("All clean.", constant.PlatePath)
}

func TestEmptySectionsHideTheirHeadings(t *testing.T) {
	o := web_interface_tester.New(t)
	o.AssertMissing("Watched Issues", constant.PlatePath)
	o.AssertMissing("Favourites", constant.PlatePath)
	o.AssertMissing("Watched Pages", constant.PlatePath)
}

func TestHiddenSectionsKeepTheirLiveTargets(t *testing.T) {
	o := web_interface_tester.New(t)
	o.AssertContains(`sse-swap="plate"`, constant.PlatePath)
	o.AssertContains(`sse-swap="watched_issues"`, constant.PlatePath)
	o.AssertContains(`sse-swap="favorites"`, constant.PlatePath)
	o.AssertContains(`sse-swap="watched_pages"`, constant.PlatePath)
}

func TestNewestSectionAbsentWithoutConfiguredProject(t *testing.T) {
	o := web_interface_tester.New(t)
	o.AssertMissing("Newest Issues", constant.PlatePath)
	o.AssertMissing(`sse-swap="newest"`, constant.PlatePath)
}

func TestNewestSectionKeepsLiveTargetWithConfiguredProject(t *testing.T) {
	o := web_interface_tester.NewWithProject(t, []string{"ABC"})
	o.AssertContains(`sse-swap="newest"`, constant.PlatePath)
	o.AssertMissing("Newest Issues", constant.PlatePath)
}
