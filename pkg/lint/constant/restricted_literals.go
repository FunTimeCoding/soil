package constant

import "github.com/funtimecoding/soil/pkg/lint/types/restriction"

var RestrictedLiterals = []restriction.Restriction{
	{
		Package: HTTPPackage,
		Symbol:  "Client",
		AllowedIn: []string{
			"github.com/funtimecoding/soil/pkg/web/stall_client/pool",
			"github.com/funtimecoding/soil/pkg/generative/model_context_client",
			"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context",
		},
		Message: OwnedClientMessage,
	},
	{
		Package: HTTPPackage,
		Symbol:  "Transport",
		AllowedIn: []string{
			"github.com/funtimecoding/soil/pkg/web/stall_client",
		},
		Message: OwnedClientMessage,
	},
}
