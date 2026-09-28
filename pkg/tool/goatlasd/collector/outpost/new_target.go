package outpost

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/face"

func NewTarget(
	name string,
	source face.OutpostSource,
) *Target {
	return &Target{Name: name, Source: source}
}
