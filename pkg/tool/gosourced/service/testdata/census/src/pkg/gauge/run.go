package gauge

import (
	"example/pkg/pair"
	"fmt"
)

func Same(n int) int {
	return pair.Compare(n, n)
}

func Different(n int, m int) int {
	fmt.Println(7)

	return pair.Compare(n, m)
}
