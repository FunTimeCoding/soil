package entry

import (
	"github.com/funtimecoding/soil/pkg/web/component/entry/field"
	"maragu.dev/gomponents"
)

type Entry struct {
	path          string
	submit        string
	identifier    string
	target        string
	swap          string
	previewPath   string
	previewTarget string
	heading       string
	alert         string
	inline        bool
	record        bool
	reset         bool
	query         bool
	class         string
	buttonClass   string
	hidden        []gomponents.Node
	alternatives  []gomponents.Node
	beside        []gomponents.Node
	fields        []*field.Field
}
