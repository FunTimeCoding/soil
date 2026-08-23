package instrument

import "github.com/funtimecoding/soil/pkg/face"

func (i *Instrument) Recorder() face.Recorder {
	return i.recorder
}
