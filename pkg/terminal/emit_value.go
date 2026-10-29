package terminal

import "github.com/funtimecoding/soil/pkg/notation"

func (t *Terminal) EmitValue(
	v any,
	e error,
) {
	if e != nil {
		t.Exitln(e)

		return
	}

	t.Line(string(notation.Marshal(v)))
}
