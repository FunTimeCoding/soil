package merge_request

import "github.com/funtimecoding/soil/pkg/face"

func (r *Request) SetAgeColor(f face.SprintFunction) {
	r.ageColor = f
}
