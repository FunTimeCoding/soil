package pattern_site

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/result/location"

type Entry struct {
	Shape    string
	Exemplar string
	Location *location.Location
}
