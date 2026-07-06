package server

import "github.com/funtimecoding/soil/pkg/face"

func (b *Builder) WithRecorder(r face.Recorder) *Builder {
	b.recorder = r

	return b
}
