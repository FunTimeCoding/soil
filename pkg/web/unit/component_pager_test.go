package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/pager"
	"github.com/funtimecoding/soil/pkg/web/web_tester"
	"testing"
)

func TestCountedPagerLinksBothWaysAroundItsPosition(t *testing.T) {
	assert.String(
		t,
		`<p class="pager"><a href="/alfa?offset=0">Previous</a><span class="pager-position">21–40 of 93</span><a href="/alfa?offset=40">Next</a></p>`,
		web_tester.New(t).Render(
			pager.New().
				Previous("/alfa?offset=0", "Previous").
				Position("21–40 of 93").
				Next("/alfa?offset=40", "Next").
				Node(),
		),
	)
}

func TestOpenEndedPagerShowsOnlyWayOnward(t *testing.T) {
	assert.String(
		t,
		`<p class="pager"><a href="/alfa?page=2">Older →</a></p>`,
		web_tester.New(t).Render(
			pager.New().Next("/alfa?page=2", "Older →").Node(),
		),
	)
}

func TestPagerWithNowhereToGoRendersNothing(t *testing.T) {
	assert.String(t, ``, web_tester.New(t).Render(pager.New().Node()))
}
