package unit

import "github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"

func netboxMachine(
	name string,
	status string,
) *client.VirtualMachine {
	result := &client.VirtualMachine{}
	result.Name = name
	result.Status = &status

	return result
}
