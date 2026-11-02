package commit

import "slices"

func (s *Service) reaching(
	directory string,
	packagePaths []string,
) []string {
	var result []string

	if len(packagePaths) == 0 {
		return nil
	}

	for _, other := range s.inventory.Replacing(directory) {
		w := s.workspace.Workspace(other)

		if slices.ContainsFunc(packagePaths, w.Reaches) {
			result = append(result, other)
		}
	}

	return result
}
