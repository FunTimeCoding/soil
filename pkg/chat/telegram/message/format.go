package message

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (m *Message) Format(f *option.Format) string {
	return status.New(f).String(m.From, m.Text).RawList(m.Raw).Format()
}
