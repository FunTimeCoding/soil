package resolver

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/url"
	"slices"
	"strings"
)

func resolveLocator(
	hosts []string,
	candidate string,
) *pointer.Resolution {
	if !strings.HasPrefix(candidate, webConstant.SecurePrefix) &&
		!strings.HasPrefix(candidate, webConstant.InsecurePrefix) {
		result := pointer.NewResolution(constant.VerdictTallied)
		result.Reason = constant.ReasonScheme

		return result
	}

	parsed, e := url.Parse(candidate)

	if e != nil {
		return pointer.NewResolution(constant.VerdictUndeclaredHost)
	}

	host := parsed.Hostname()

	if slices.Contains(constant.ImplicitHosts, host) ||
		slices.Contains(hosts, host) {
		return pointer.NewResolution(constant.VerdictLive)
	}

	return pointer.NewResolution(constant.VerdictUndeclaredHost)
}
