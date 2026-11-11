package bravo

import (
	alfaConstant "example/pkg/alfa/constant"
	"example/pkg/xray/constant"
)

func Clamp(name string) int {
	if name == alfaConstant.Prefix {
		return constant.Limit
	}

	return 0
}
