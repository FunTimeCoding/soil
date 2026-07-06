package merge_request

import "github.com/funtimecoding/soil/pkg/face"

func (r *Request) AgeColor() face.SprintFunction {
	return r.ageColor
}
