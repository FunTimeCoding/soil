package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	"net/http"
	"testing"
)

func TestSystemsPageSelectsAndKeepsSelection(t *testing.T) {
	s := base.New(t)
	s.SecurityTracker.SeedKnown()
	s.Started()
	c := s.SecurityClient()
	x := t.Context()
	candidate := "8.6.1-1"

	for name, labels := range map[string]map[string]string{
		"alfa":    {"exposure": "internet", "tenant": "one"},
		"bravo":   {"exposure": "internet", "tenant": "two"},
		"charlie": {"exposure": "internal", "tenant": "one"},
	} {
		assert.FatalOnError(
			t,
			c.PutSystem(
				x,
				name,
				client.SystemSnapshot{
					Release: "trixie",
					Labels:  &labels,
					Packages: []*client.InstalledPackage{
						{
							Name:          "libtango5",
							Version:       "8.4.0-1",
							Source:        "tango-utils",
							SourceVersion: "8.4.0-1",
							Candidate:     &candidate,
						},
					},
				},
			),
		)
	}

	v := s.Served()
	page := v.Body(t, "/systems?label=exposure%3Ainternet")
	assert.StringContains(t, "/system?name=alfa", page)
	assert.StringContains(t, "/system?name=bravo", page)
	assert.StringNotContains(t, "/system?name=charlie", page)
	assert.StringContains(
		t,
		"<strong>2</strong><br><small>Systems reporting",
		page,
	)
	assert.StringContains(
		t,
		`href="/systems?label=exposure%3Ainternet" class="chip active" aria-current="true">internet 2</a>`,
		page,
	)
	assert.StringContains(
		t,
		`href="/systems?label=exposure%3Ainternal" class="chip">internal 1</a>`,
		page,
	)
	assert.StringContains(
		t,
		`href="/systems?label=exposure%3Ainternet&amp;label=tenant%3Aone" class="chip">one 1</a>`,
		page,
	)
	assert.StringContains(
		t,
		`href="/systems?label=exposure%3Ainternet&amp;show=supported" class="chip">Supported 2</a>`,
		page,
	)
	assert.StringContains(
		t,
		`<input type="hidden" name="label" value="exposure:internet">`,
		page,
	)
	found := v.Body(t, "/systems?search=RAV")
	assert.StringContains(t, "/system?name=bravo", found)
	assert.StringNotContains(t, "/system?name=alfa", found)
	assert.StringContains(t, `value="RAV"`, found)
	assert.Integer(
		t,
		http.StatusBadRequest,
		v.Status(t, "/systems?label=exposure"),
	)
	assert.Integer(
		t,
		http.StatusBadRequest,
		v.Status(t, "/systems?label=exposure%3A"),
	)
	assert.Integer(
		t,
		http.StatusBadRequest,
		v.Status(t, "/systems?label=%3Ainternet"),
	)
}
