package main

import (
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(true)
	lorem := split.Space(
		"One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin. He lay on his armour-like back, and if he lifted his head a little he could see his brown belly, slightly domed and divided by arches into stiff sections. The bedding was hardly able to cover it and seemed ready to slide off any moment. His many legs, pitifully thin compared with the size of the rest of him, waved about helplessly as he looked. ",
	)
	columns, rows := 10, 40
	word := 0

	for r := range rows {
		for c := range columns {
			color := tcell.ColorWhite

			if c < 1 || r < 1 {
				color = tcell.ColorYellow
			}

			table.SetCell(
				r,
				c,
				tview.NewTableCell(lorem[word]).
					SetTextColor(color).
					SetAlign(tview.AlignCenter),
			)
			word = (word + 1) % len(lorem)
		}
	}

	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(
		func(key tcell.Key) {
			if key == tcell.KeyEscape {
				app.Stop()
			}

			if key == tcell.KeyEnter {
				table.SetSelectable(true, true)
			}
		},
	).SetSelectedFunc(
		func(
			row int,
			column int,
		) {
			table.GetCell(row, column).SetTextColor(tcell.ColorRed)
			table.SetSelectable(false, false)
		},
	)

	if e := app.SetRoot(table, true).EnableMouse(true).Run(); e != nil {
		panic(e)
	}
}
