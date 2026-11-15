package technitium_tester

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/technitium/response"
)

func ZoneOptions(answer string) *response.ZoneOptions {
	var result response.ZoneOptions
	notation.MustDecode(answer, &result, false)

	return &result
}
