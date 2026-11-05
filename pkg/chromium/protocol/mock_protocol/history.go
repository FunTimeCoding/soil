package mock_protocol

import "github.com/funtimecoding/soil/pkg/chromium/history"

func (p *Protocol) History() (*history.Result, error) {
	return nil, nil
}
