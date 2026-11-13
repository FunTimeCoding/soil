package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/plate"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"testing"
)

func TestPlateStacksNameHeadlineCaptionAndRows(t *testing.T) {
	assert.String(
		t,
		`<div class="plate"><div class="plate-name">alfa</div><div class="plate-headline">12 kg</div><div class="plate-caption">bravo</div><div class="plate-row"><span class="plate-label">charlie</span><div class="plate-value">delta</div></div></div>`,
		renderedNode(
			t,
			plate.New().
				Name("alfa").
				Headline("12 kg").
				Caption("bravo").
				Row("charlie", "delta").
				Node(),
		),
	)
}

func TestPlateLeavesOutEmptyLines(t *testing.T) {
	assert.String(
		t,
		`<div class="plate"><div class="plate-row"><span class="plate-label">charlie</span><div class="plate-value">delta</div></div></div>`,
		renderedNode(t, plate.New().Row("charlie", "delta").Node()),
	)
}

func TestPlateRowNodeCarriesAttributesToValue(t *testing.T) {
	assert.String(
		t,
		`<div class="plate"><div class="plate-row"><span class="plate-label">charlie</span><div class="plate-value" style="--alfa: 1"><b>delta</b></div></div></div>`,
		renderedNode(
			t,
			plate.New().
				RowNode(
					"charlie",
					gomponents.Group{
						html.Style("--alfa: 1"),
						html.B(gomponents.Text("delta")),
					},
				).
				Node(),
		),
	)
}

func TestPlateGroupWrapsPlates(t *testing.T) {
	assert.String(
		t,
		`<div class="plate-group"><div class="plate"><div class="plate-headline">1 kg</div></div><div class="plate"><div class="plate-headline">2 kg</div></div></div>`,
		renderedNode(
			t,
			plate.Group(
				plate.New().Headline("1 kg").Node(),
				plate.New().Headline("2 kg").Node(),
			),
		),
	)
}
