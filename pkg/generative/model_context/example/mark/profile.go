package mark

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func profile(identifier string) (string, error) {
	if identifier == "123" {
		return `{"id": "123", "name": "John Doe", "email": "john@gmail.com"}`, nil
	}

	return "", not_found.Format("user not found")
}
