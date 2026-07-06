package run

import "github.com/funtimecoding/soil/pkg/strings/join/key_value"

func (r *Run) Environment(
	k string,
	v string,
) {
	r.environment = append(r.environment, key_value.Equals(k, v))
}
