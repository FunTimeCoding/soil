package unit

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/technitium/response"
)

func settings(answer string) *response.Settings {
	var result response.Settings
	notation.MustDecode(answer, &result, false)

	return &result
}
