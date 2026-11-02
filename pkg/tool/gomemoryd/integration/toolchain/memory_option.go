package toolchain

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"

func memoryOption(
	name string,
	content string,
) *save_option.Option {
	result := save_option.New()
	result.Name = name
	result.Content = content
	result.Description = name
	result.Source = "test"

	return result
}
