package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument/update_machine"
)

func convertUpdateMachine(
	r server.UpdateMachineRequestObject,
) *update_machine.Machine {
	m := update_machine.New()
	m.Identifier = int(r.Identifier)

	if r.Params.Node != nil {
		m.Node = *r.Params.Node
	}

	if r.Body.Name != nil {
		m.Name = *r.Body.Name
	}

	if r.Body.Tags != nil {
		m.Tags = *r.Body.Tags
	}

	if r.Body.Onboot != nil {
		m.OnBoot = r.Body.Onboot
	}

	if r.Body.Cores != nil {
		m.Cores = *r.Body.Cores
	}

	if r.Body.Memory != nil {
		m.Memory = *r.Body.Memory
	}

	if r.Body.Description != nil {
		m.Description = *r.Body.Description
	}

	if r.Body.Delete != nil {
		m.Delete = *r.Body.Delete
	}

	return m
}
