package console

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/face"
)

func Gradient(
	startHex string,
	endHex string,
	n int,
) []face.SprintFunction {
	r1, g1, b1 := parseHex(startHex)
	r2, g2, b2 := parseHex(endHex)
	result := make([]face.SprintFunction, n)

	for i := 0; i < n; i++ {
		r := int(
			float64(r1) +
				float64(i)*(float64(r2)-float64(r1))/float64(n-1),
		)
		g := int(
			float64(g1) +
				float64(i)*(float64(g2)-float64(g1))/float64(n-1),
		)
		b := int(
			float64(b1) +
				float64(i)*(float64(b2)-float64(b1))/float64(n-1),
		)
		hex := fmt.Sprintf("#%02x%02x%02x", r, g, b)
		result[i] = NewColor(hex)
	}

	return result
}
