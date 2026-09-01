package build

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
	"os"
)

func Headers(token string) map[string]string {
	headers := make(map[string]string)

	if token != "" {
		headerName, headerValue := key_value.Equals(token)
		console.Format("Header name: %s\n", headerName)
		console.Format("Header value: %s\n", headerValue)

		if headerValue == "" {
			console.Line("Header value empty")
			os.Exit(1)
		}

		headers[headerName] = headerValue
	}

	return headers
}
