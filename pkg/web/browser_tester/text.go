package browser_tester

import "fmt"

func (b *Browser) Text(selector string) string {
	b.T.Helper()
	var result string
	b.Evaluate(
		fmt.Sprintf(
			"(document.querySelector('%s') || {}).textContent || ''",
			selector,
		),
		&result,
	)

	return result
}
