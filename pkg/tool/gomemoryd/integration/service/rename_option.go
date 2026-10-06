package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
)

func renameOption(
	m *record.Memory,
	name string,
) *save_option.Option {
	o := save_option.New()
	o.Name = name
	o.Content = m.Content
	o.Description = m.Description
	o.Source = "test"

	return o
}
