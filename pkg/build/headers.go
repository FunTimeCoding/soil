package build

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
	"os"
)

func Headers(token string) map[string]string {
	headers := make(map[string]string)

	if token != "" {
		headerName, headerValue := key_value.Equals(token)
		fmt.Printf("Header name: %s\n", headerName)
		fmt.Printf("Header value: %s\n", headerValue)

		if headerValue == "" {
			fmt.Println("Header value empty")
			os.Exit(1)
		}

		headers[headerName] = headerValue
	}

	return headers
}
