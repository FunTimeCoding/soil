package plate

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func line(
	class string,
	text string,
) gomponents.Node {
	if text == "" {
		return nil
	}

	return html.Div(html.Class(class), gomponents.Text(text))
}
