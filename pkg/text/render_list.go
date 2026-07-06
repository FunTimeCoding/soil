package text

import "github.com/funtimecoding/soil/pkg/text/multi_line"

func RenderList(
	title string,
	v []string,
) string {
	l := multi_line.New()
	l.Format("%s:", title)

	for _, e := range v {
		l.Format("%s", e)
	}

	return l.Render()
}
