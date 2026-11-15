package client_tester

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func OperationMethod(identifier string) string {
	result := join.Empty(strings.ToUpper(identifier[:1]), identifier[1:])

	if result == "GetAnsibleInventory" {
		return "AnsibleInventoryMarkup"
	}

	return result
}
