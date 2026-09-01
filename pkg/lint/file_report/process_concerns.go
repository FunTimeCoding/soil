package file_report

import "github.com/funtimecoding/soil/pkg/console"

func (r *Report) ProcessConcerns(fix bool) {
	if r.HasConcerns() {
		if fix && r.Fix != nil {
			r.Fix()
		} else {
			for _, c := range r.Concerns {
				console.Format("%s: %s\n", c.Text, c.Path)
			}
		}
	}
}
