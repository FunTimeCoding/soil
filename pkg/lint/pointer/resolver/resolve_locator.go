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
		return pointer.NewUnchecked(constant.ReasonScheme)
	}

	parsed, e := url.Parse(candidate)

	if e != nil {
		return pointer.NewConcern(constant.UndeclaredHostKey)
	}

	host := parsed.Hostname()

	if slices.Contains(constant.ImplicitHosts, host) ||
		slices.Contains(hosts, host) {
		return pointer.NewLive("")
	}

	return pointer.NewConcern(constant.UndeclaredHostKey)
}
