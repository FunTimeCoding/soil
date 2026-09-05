package browser_tester

import "fmt"

func (b *Browser) CountElements(selector string) int {
	b.T.Helper()
	var count int
	b.Evaluate(
		fmt.Sprintf("document.querySelectorAll('%s').length", selector),
		&count,
	)

	return count
}
