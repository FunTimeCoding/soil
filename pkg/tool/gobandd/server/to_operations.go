package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/types/operation"
)

func toOperations(v []*operation.Operation) []server.Operation {
	result := make([]server.Operation, 0, len(v))

	for _, o := range v {
		element := server.Operation{
			Name:  server.OperationName(o.Name),
			Force: o.Force,
			State: o.State,
		}

		if o.Refusal != "" {
			element.Refusal = new(o.Refusal)
		}

		result = append(result, element)
	}

	return result
}
