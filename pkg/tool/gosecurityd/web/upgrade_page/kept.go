package upgrade_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"net/url"
	"strconv"
)

func Kept(
	release string,
	kernel bool,
) url.Values {
	result := url.Values{}

	if release != "" {
		result.Set(constant.ReleaseParameter, release)
	}

	if kernel {
		result.Set(constant.KernelParameter, strconv.FormatBool(kernel))
	}

	return result
}
