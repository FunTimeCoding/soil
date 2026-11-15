package measure_tester

import "github.com/funtimecoding/soil/pkg/measure/count"

func Counted(
	code int,
	comment int,
	blank int,
) *count.Count {
	c := count.New()
	c.Code = code
	c.Comment = comment
	c.Blank = blank

	return c
}
