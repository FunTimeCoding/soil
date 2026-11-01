package field

import "github.com/funtimecoding/soil/pkg/web/component/choice"

func (f *Field) Choices(choices []*choice.Choice) *Field {
	f.choices = choices

	return f
}
