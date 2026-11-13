package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/panel"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"testing"
)

func TestPanelIsTitledArticle(t *testing.T) {
	assert.String(
		t,
		`<article id="alfa"><header><strong>Bravo</strong></header><p>charlie</p></article>`,
		renderedNode(
			t,
			panel.New("Bravo").
				Identifier("alfa").
				Content(html.P(gomponents.Text("charlie"))).
				Node(),
		),
	)
}

func TestPanelFramesItsHeaderAndFooter(t *testing.T) {
	assert.String(
		t,
		`<article><header>delta</header><p>charlie</p><footer>echo</footer></article>`,
		renderedNode(
			t,
			panel.New("").
				Header(gomponents.Text("delta")).
				Content(html.P(gomponents.Text("charlie"))).
				Footer(gomponents.Text(constant.FixtureEcho)).
				Node(),
		),
	)
}

func TestPanelTitleSharesBandWithHeaderInRow(t *testing.T) {
	assert.String(
		t,
		`<article><header class="heading-row"><strong>Bravo</strong>delta</header><p>charlie</p></article>`,
		renderedNode(
			t,
			panel.New("Bravo").
				Header(gomponents.Text("delta")).
				Content(html.P(gomponents.Text("charlie"))).
				Node(),
		),
	)
}
