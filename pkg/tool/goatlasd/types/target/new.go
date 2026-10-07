package target

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/face"

func New(
	name string,
	source face.OutpostSource,
) *Target {
	return &Target{Name: name, Source: source}
}
