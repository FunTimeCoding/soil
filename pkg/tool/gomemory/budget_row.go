package gomemory

import "fmt"

func budgetRow(
	name string,
	tokens int,
	shown int,
	trimmed int,
) string {
	return fmt.Sprintf("%-13s %7d %7d %8d", name, tokens, shown, trimmed)
}
