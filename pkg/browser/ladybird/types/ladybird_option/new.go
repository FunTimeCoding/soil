package ladybird_option

import "github.com/funtimecoding/soil/pkg/browser/constant"

func New() *Option {
	return &Option{DriverPath: constant.LadybirdDriverPath, Headless: true}
}
