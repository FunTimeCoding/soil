package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/url"
	"slices"
	"strings"
)

func resolveLocator(
	hosts []string,
	candidate string,
) *Resolution {
	if !strings.HasPrefix(candidate, webConstant.SecurePrefix) &&
		!strings.HasPrefix(candidate, webConstant.InsecurePrefix) {
		return &Resolution{
			Verdict: constant.VerdictTallied,
			Reason:  constant.ReasonScheme,
		}
	}

	parsed, e := url.Parse(candidate)

	if e != nil {
		return &Resolution{Verdict: constant.VerdictUndeclaredHost}
	}

	host := parsed.Hostname()

	if slices.Contains(constant.ImplicitHosts, host) ||
		slices.Contains(hosts, host) {
		return &Resolution{Verdict: constant.VerdictLive}
	}

	return &Resolution{Verdict: constant.VerdictUndeclaredHost}
}
