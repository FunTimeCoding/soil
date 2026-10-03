package console

import "github.com/muesli/termenv"

// Reference: https://iterm2.com/feature-reporting/Hyperlinks_in_Terminal_Emulators.html
func Link(
	link string,
	text string,
	osc8 bool,
) string {
	if osc8 {
		return termenv.Hyperlink(link, text)
	}

	return link
}
