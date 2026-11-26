package request

import "github.com/funtimecoding/soil/pkg/browser/constant"

func NewText(value string) *Text {
	return &Text{Type: constant.StringValue, Value: value}
}
