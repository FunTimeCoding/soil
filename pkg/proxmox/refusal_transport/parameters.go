package refusal_transport

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/proxmox/refusal_transport/parameter_errors"
	"io"
	"maps"
	"slices"
	"strings"
)

func parameters(body io.Reader) string {
	var payload parameter_errors.ParameterErrors

	if json.NewDecoder(body).Decode(&payload) != nil ||
		len(payload.Errors) == 0 {
		return ""
	}

	var result []string

	for _, k := range slices.Sorted(maps.Keys(payload.Errors)) {
		result = append(result, fmt.Sprintf("%s: %s", k, payload.Errors[k]))
	}

	return fmt.Sprintf(" - %s", strings.Join(result, "; "))
}
