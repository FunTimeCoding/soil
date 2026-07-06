package locator

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"github.com/funtimecoding/soil/pkg/system/join"
	"net"
	"net/url"
	"strconv"
	"strings"
)

func (l *Locator) String() string {
	o := &url.URL{Scheme: l.scheme}

	if l.host != "" {
		if l.port != 0 {
			o.Host = net.JoinHostPort(l.host, strconv.Itoa(l.port))
		} else {
			o.Host = l.host
		}
	}

	if l.basePath != "" || l.path != "" {
		o.Path = join.Absolute(l.basePath, l.path)
	}

	if l.trail && !strings.HasSuffix(o.Path, separator.Slash) {
		o.Path = fmt.Sprintf("%s%s", o.Path, separator.Slash)
	}

	o.RawQuery = l.value.Encode()

	if l.user != "" && l.password == "" {
		o.User = url.User(l.user)
	} else if l.user != "" && l.password != "" {
		o.User = url.UserPassword(l.user, l.password)
	}

	result := o.String()

	if l.fragment != "" {
		if len(l.fragmentValue) > 0 {
			result = fmt.Sprintf(
				"%s#%s?%s",
				result,
				l.fragment,
				l.fragmentValue.Encode(),
			)
		} else {
			result = fmt.Sprintf("%s#%s", result, l.fragment)
		}
	}

	return result
}
