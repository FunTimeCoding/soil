package reporter

import "github.com/fatih/color"

func (r *Reporter) paint(
	a color.Attribute,
	s string,
) string {
	if !r.color {
		return s
	}

	c := color.New(a)
	c.EnableColor()

	return c.Sprint(s)
}
