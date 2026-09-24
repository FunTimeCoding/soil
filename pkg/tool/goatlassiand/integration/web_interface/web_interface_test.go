package web_interface

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/integration/web_interface_tester"
	"testing"
)

func TestPlatePageRendersBothIssueSections(t *testing.T) {
	o := web_interface_tester.New(t)
	o.AssertContains("Plate", constant.PlatePath)
	o.AssertContains("Watched Issues", constant.PlatePath)
}

func TestPlatePageRendersBothPageSections(t *testing.T) {
	o := web_interface_tester.New(t)
	o.AssertContains("Favourites", constant.PlatePath)
	o.AssertContains("Watched Pages", constant.PlatePath)
}

func TestWatchedIssuesSectionIsLive(t *testing.T) {
	o := web_interface_tester.New(t)
	o.AssertContains(`sse-swap="watched_issues"`, constant.PlatePath)
}

func TestEmptyIssueSectionsNameWhichIsEmpty(t *testing.T) {
	o := web_interface_tester.New(t)
	o.AssertContains("Nothing on the plate.", constant.PlatePath)
	o.AssertContains("No watched issues.", constant.PlatePath)
}

func TestNewestSectionAbsentWithoutConfiguredProject(t *testing.T) {
	o := web_interface_tester.New(t)
	o.AssertMissing("Newest Issues", constant.PlatePath)
	o.AssertMissing(`sse-swap="newest"`, constant.PlatePath)
}

func TestNewestSectionPresentWithConfiguredProject(t *testing.T) {
	o := web_interface_tester.NewWithProject(t, []string{"ABC"})
	o.AssertContains("Newest Issues", constant.PlatePath)
	o.AssertContains(`sse-swap="newest"`, constant.PlatePath)
	o.AssertContains("No recent issues.", constant.PlatePath)
}
