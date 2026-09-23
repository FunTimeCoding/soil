package web_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/chart"
	"strings"
)

func RenderChart(c *chart.Chart) string {
	var b strings.Builder
	errors.PanicOnError(c.Render().Render(&b))

	return b.String()
}
