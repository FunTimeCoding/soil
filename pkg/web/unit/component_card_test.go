package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/card"
	"github.com/funtimecoding/soil/pkg/web/web_tester"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"testing"
)

func TestBareCardIsNamedArticle(t *testing.T) {
	assert.String(
		t,
		`<article class="card"><div class="card-body"><header class="card-head"><h4>Alfa</h4></header></div></article>`,
		web_tester.New(t).Render(card.New("Alfa").Node()),
	)
}

func TestCardFillsEverySlotItIsGiven(t *testing.T) {
	assert.String(
		t,
		`<article id="card-bravo" class="card card-side card-hero" style="--card-accent: #4a7fb5" data-search="bravo worker"><div class="card-visual"><img src="/bravo.png" alt="Bravo" loading="lazy"></div><div class="card-body"><header class="card-head"><i>dot</i><h1><a class="card-link" href="/bravo">Bravo</a></h1><b>canary</b></header><p class="card-subtitle">worker</p><p class="card-meta">staging</p><p>state</p><footer class="card-footer"><button>restart</button></footer></div></article>`,
		web_tester.New(t).Render(
			card.New("Bravo").
				Identifier("card-bravo").
				Link("/bravo").
				Portrait("/bravo.png").
				Emblem("⚙").
				Side().
				Hero().
				Accent("#4a7fb5").
				Marker(html.I(gomponents.Text("dot"))).
				Trailing(html.B(gomponents.Text("canary"))).
				Subtitle("worker").
				Meta("staging").
				State(html.P(gomponents.Text("state"))).
				Footer(html.Button(gomponents.Text("restart"))).
				Terms("bravo worker").
				Node(),
		),
	)
}

func TestCardWithoutPortraitFallsBackToItsEmblem(t *testing.T) {
	assert.StringContains(
		t,
		`<article class="card card-art"><div class="card-visual card-fallback">⚙</div>`,
		web_tester.New(t).Render(card.New("Charlie").Emblem("⚙").Art().Node()),
	)
}

func TestTaskCardMarksItselfBesideItsSize(t *testing.T) {
	assert.StringContains(
		t,
		`<article class="card card-compact card-task extra">`,
		web_tester.New(t).Render(
			card.New("Delta").Compact().Task().Class("extra").Node(),
		),
	)
}
