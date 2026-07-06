package face

import (
	"github.com/funtimecoding/soil/pkg/console/description"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

type Formattable interface {
	Format(f *option.Format) string
	Meta() *description.Description
}
