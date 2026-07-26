package service

import "os"

func captureSourceBytes(plan *movePlan) (map[string][]byte, error) {
	result := make(map[string][]byte)

	for _, entry := range plan.entries {
		filename := plan.set.Position(entry.file.Pos()).Filename

		if _, exists := result[filename]; exists {
			continue
		}

		content, e := os.ReadFile(filename)

		if e != nil {
			return nil, e
		}

		result[filename] = content
	}

	return result, nil
}
