package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/panel"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"testing"
)

func TestAPanelIsATitledArticle(t *testing.T) {
	assert.String(
		t,
		`<article id="alfa"><h4>Bravo</h4><p>charlie</p></article>`,
		renderedNode(
			t,
			panel.New("Bravo").
				Identifier("alfa").
				Content(html.P(gomponents.Text("charlie"))).
				Node(),
		),
	)
}

func TestAPanelFramesItsHeaderAndFooter(t *testing.T) {
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
