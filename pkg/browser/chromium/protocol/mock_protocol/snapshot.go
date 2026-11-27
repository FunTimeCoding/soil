package mock_protocol

import "github.com/funtimecoding/soil/pkg/browser/chromium/snapshot"

func (p *Protocol) Snapshot() ([]*snapshot.Node, error) {
	return nil, nil
}
