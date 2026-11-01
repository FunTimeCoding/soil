package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/memory_list"
	"testing"
)

func TestRemovingOneMemoryFilterKeepsTheOthers(t *testing.T) {
	l := memory_list.New("feedback", "shared", "work", 1, 20)
	page := renderList(t, l)
	assert.StringContains(
		t,
		`<a href="/memories?scope=work&amp;type=feedback" class="chip">tag: shared ×</a>`,
		page,
	)
	assert.StringContains(
		t,
		`<a href="/memories?scope=work&amp;tag=shared" class="chip">type: feedback ×</a>`,
		page,
	)
}

func TestMemoryPagerCarriesTheFilters(t *testing.T) {
	l := memory_list.New("", "shared", "", 2, 20)
	l.Total = 45
	page := renderList(t, l)
	assert.StringContains(t, "21–40 of 45", page)
	assert.StringContains(t, `href="/memories?tag=shared"`, page)
	assert.StringContains(t, `href="/memories?page=3&amp;tag=shared"`, page)
}

func TestTheChosenScopeIsCurrent(t *testing.T) {
	l := memory_list.New("", "", "work", 1, 20)
	l.Scopes = []string{"", "work"}
	page := renderList(t, l)
	assert.StringContains(
		t,
		`<a href="/memories?scope=work" class="chip active" aria-current="true">work</a>`,
		page,
	)
	assert.StringContains(t, `<a href="/memories" class="chip">all</a>`, page)
}
