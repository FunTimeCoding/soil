package save_tester

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"

func Update() *save_option.Option {
	o := save_option.New()
	o.Name = "based entry"
	o.Content = "updated"
	o.Description = "based entry"
	o.Source = "test"

	return o
}
