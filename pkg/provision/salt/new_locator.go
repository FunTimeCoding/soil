package salt

import "github.com/funtimecoding/soil/pkg/web/locator"

func newLocator(
	host string,
	port int,
	insecure bool,
) *locator.Locator {
	result := locator.New(host)

	if port != 0 {
		result.Port(port)
	}

	if insecure {
		result.Insecure()
	}

	return result
}
