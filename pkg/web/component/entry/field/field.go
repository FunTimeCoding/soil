package field

import (
	"github.com/funtimecoding/soil/pkg/web/component/choice"
	"maragu.dev/gomponents"
)

type Field struct {
	kind         string
	name         string
	label        string
	value        string
	placeholder  string
	prompt       string
	required     bool
	focus        bool
	checked      bool
	minimum      string
	maximum      string
	step         string
	lines        string
	class        string
	accept       string
	revealPath   string
	revealTarget string
	suggestions  []string
	choices      []*choice.Choice
	content      []gomponents.Node
}
