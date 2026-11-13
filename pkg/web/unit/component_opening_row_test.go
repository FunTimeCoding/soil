package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/opening_row"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"testing"
)

func TestClosedRowAsksToOpenAboveHiddenPlace(t *testing.T) {
	assert.String(
		t,
		`<tr id="alfa" class="opening-row" aria-expanded="false" hx-get="/open" hx-target="#alfa-detail" hx-swap="outerHTML"><td>bravo</td></tr><tr id="alfa-detail" hidden="hidden"></tr>`,
		renderedNode(
			t,
			opening_row.New("alfa", "/open", "/close").
				Cells(html.Td(gomponents.Text("bravo"))).
				Node(),
		),
	)
}

func TestOpenRowAsksToCloseAboveItsDetail(t *testing.T) {
	assert.String(
		t,
		`<tr id="alfa" class="opening-row" aria-expanded="true" hx-get="/close" hx-target="#alfa-detail" hx-swap="outerHTML"><td>bravo</td></tr><tr id="alfa-detail" class="opening-detail"><td colspan="3"><div class="opening-content">charlie</div></td></tr>`,
		renderedNode(
			t,
			opening_row.New("alfa", "/open", "/close").
				Cells(html.Td(gomponents.Text("bravo"))).
				Detail(3, gomponents.Text("charlie")).
				Node(),
		),
	)
}

func TestResponseSendsRowOutOfBandBesideDetail(t *testing.T) {
	assert.String(
		t,
		`<tr id="alfa-detail" class="opening-detail"><td colspan="3"><div class="opening-content">charlie</div></td></tr><template><tr id="alfa" class="opening-row" aria-expanded="true" hx-get="/close" hx-target="#alfa-detail" hx-swap="outerHTML" hx-swap-oob="true"><td>bravo</td></tr></template>`,
		renderedNode(
			t,
			opening_row.New("alfa", "/open", "/close").
				Cells(html.Td(gomponents.Text("bravo"))).
				Detail(3, gomponents.Text("charlie")).
				Response(),
		),
	)
}

func TestRemovalDeletesRowOutOfBand(t *testing.T) {
	assert.String(
		t,
		`<template><tr id="alfa" hx-swap-oob="delete"></tr></template>`,
		renderedNode(t, opening_row.New("alfa", "/open", "/close").Removal()),
	)
}
