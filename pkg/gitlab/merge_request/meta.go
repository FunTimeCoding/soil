package merge_request

import "github.com/funtimecoding/soil/pkg/console/description"

func (r *Request) Meta() *description.Description {
	return description.New("Merge request", "Merge request")
}
