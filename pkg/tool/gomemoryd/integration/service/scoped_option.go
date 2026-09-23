package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
)

func scopedOption(
	name string,
	scope string,
) *save_option.Option {
	o := save_option.New()
	o.Name = name
	o.Content = constant.FixtureContent
	o.Description = name
	o.Source = "test"
	o.Scope = scope

	return o
}
