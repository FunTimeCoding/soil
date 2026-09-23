package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"strings"
)

func frameNames(d *decision.Decision) string {
	var result []string

	for _, v := range d.Frames {
		result = append(result, v.Name)
	}

	return strings.Join(result, " · ")
}
