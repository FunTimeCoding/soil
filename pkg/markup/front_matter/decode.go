package front_matter

import "github.com/funtimecoding/soil/pkg/markup"

func (f *Front) Decode(structure any) error {
	return markup.Decode(f.Raw, structure)
}
