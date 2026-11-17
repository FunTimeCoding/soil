package tool_tester

import "github.com/funtimecoding/soil/pkg/notation"

func NewSchema(input any) *Schema {
	result := &Schema{}
	notation.MustDecode(string(notation.Marshal(input)), result, false)

	return result
}
