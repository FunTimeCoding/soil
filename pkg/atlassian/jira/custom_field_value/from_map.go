package custom_field_value

import "github.com/funtimecoding/soil/pkg/notation"

func FromMap(m map[string]any) *Value {
	var result Value
	notation.MustDecode(
		notation.Encode(m, false),
		&result,
		true,
	)

	return &result
}
