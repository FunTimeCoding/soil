package unit

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/technitium/response"
)

func zoneOptions(answer string) *response.ZoneOptions {
	var result response.ZoneOptions
	notation.MustDecode(answer, &result, false)

	return &result
}
