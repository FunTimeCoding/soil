package mock_page

import "github.com/funtimecoding/soil/pkg/chromium/history"

func (p *Page) History() (*history.Result, error) {
	return nil, nil
}
