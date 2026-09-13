package browser_tester

import "fmt"

func (b *Browser) Style(
	selector string,
	property string,
) string {
	b.T.Helper()
	var result string
	b.Evaluate(
		fmt.Sprintf(
			"getComputedStyle(document.querySelector('%s')).%s",
			selector,
			property,
		),
		&result,
	)

	return result
}
