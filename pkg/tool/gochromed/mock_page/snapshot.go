package mock_page

import "github.com/funtimecoding/soil/pkg/chromium/snapshot"

func (p *Page) Snapshot() ([]*snapshot.Node, error) {
	return nil, nil
}
