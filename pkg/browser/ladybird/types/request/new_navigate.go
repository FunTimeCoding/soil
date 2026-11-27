package request

import "github.com/funtimecoding/soil/pkg/browser/constant"

func NewNavigate(
	context string,
	locator string,
) *Navigate {
	return &Navigate{
		Context: context,
		Locator: locator,
		Wait:    constant.LadybirdCompleteReadiness,
	}
}
