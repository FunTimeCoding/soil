package comparison

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/section"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/section/parser"
)

func wordCounts(body string) (map[string]map[string]int, []string) {
	result := map[string]map[string]int{}
	var order []string

	for _, x := range parser.Parse(body) {
		if result[x.Title] == nil {
			result[x.Title] = map[string]int{}
			order = append(order, x.Title)
		}

		for _, w := range section.Words(x) {
			result[x.Title][w]++
		}
	}

	return result, order
}
