package fact

import "maragu.dev/gomponents"

func (f *Fact) Node() gomponents.Node {
	if f.glance {
		return f.glanceNode()
	}

	return f.table()
}
