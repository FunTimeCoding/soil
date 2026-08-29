package mock_service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument/update_machine"

func (s *Service) UpdateMachine(
	_ string,
	a *update_machine.Machine,
) error {
	return a.Validate()
}
