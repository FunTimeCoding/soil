package mock_service

import "fmt"

func (s *Service) ResolveInstance(explicit string) (string, error) {
	if explicit != "" {
		if _, okay := s.Instance(explicit); okay {
			return explicit, nil
		}

		return "", fmt.Errorf("unknown instance: %s", explicit)
	}

	if len(s.inventory.Instances) == 1 {
		return s.inventory.Instances[0].Name, nil
	}

	return "", fmt.Errorf(
		"no instance selected - %d instances configured, selection required",
		len(s.inventory.Instances),
	)
}
