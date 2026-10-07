package gw2

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gw2/log_manager/response"
	"github.com/funtimecoding/soil/pkg/notation"
)

func ParseGuilds(s string) []*response.Guild {
	var guilds map[string]any
	notation.MustDecode(s, &guilds, false)
	var result []*response.Guild

	for k, v := range guilds {
		if v == nil {
			errors.Warning("no data: %s", k)

			continue
		}

		var g response.Guild

		if e := notation.Decode(notation.Encode(v, false), &g); e != nil {
			errors.PanicOnError(e)
		}

		result = append(result, &g)
	}

	return result
}
