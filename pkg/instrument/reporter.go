package instrument

import "github.com/funtimecoding/soil/pkg/face"

func (i *Instrument) Reporter() face.Reporter {
	return i.reporter
}
