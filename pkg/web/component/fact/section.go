package fact

import "github.com/funtimecoding/soil/pkg/web/component/fact/row"

func (f *Fact) Section(label string) *Fact {
	r := row.New(label, nil)
	r.Section = true
	f.rows = append(f.rows, r)

	return f
}
