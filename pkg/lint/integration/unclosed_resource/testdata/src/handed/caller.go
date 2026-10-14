package example

import "example/helper"

func NewHolder() *helper.Holder {
	return &helper.Holder{Resource: helper.Open()}
}

func NewLeaky() *helper.Leaky {
	return &helper.Leaky{Resource: helper.Open()}
}

func UseShared() string {
	r := helper.Shared()

	return r.Name()
}
