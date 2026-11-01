package dial

import "maragu.dev/gomponents"

type Dial struct {
	track     string
	fillColor string
	fillShare float64
	filled    bool
	layers    []*layer
	markers   []*marker
	center    []gomponents.Node
}
