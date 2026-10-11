package unit

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"

func updateOption() *save_option.Option {
	o := save_option.New()
	o.Name = "based entry"
	o.Content = "updated"
	o.Description = "based entry"
	o.Source = "test"

	return o
}
