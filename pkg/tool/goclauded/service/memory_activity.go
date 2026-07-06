package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/service/memory_activity"

func (s *Service) MemoryActivity(since string) []memory_activity.Activity {
	versions := s.memory.VersionsSince(since, 50)
	var result []memory_activity.Activity

	for _, v := range versions {
		result = append(
			result,
			memory_activity.Activity{
				MemoryIdentifier: v.MemoryIdentifier,
				Name:             v.Name,
				ChangeType:       v.ChangeType,
				Source:           v.Source,
			},
		)
	}

	return result
}
