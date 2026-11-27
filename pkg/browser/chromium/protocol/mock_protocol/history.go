package mock_protocol

import "github.com/funtimecoding/soil/pkg/browser/chromium/history"

func (p *Protocol) History() (*history.Result, error) {
	return nil, nil
}
