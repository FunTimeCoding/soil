package sweep

import "path/filepath"

func collectSources(directories []string) []string {
	var result []string

	for _, d := range directories {
		files, e := filepath.Glob(filepath.Join(d, "*.jsonl"))

		if e != nil {
			continue
		}

		result = append(result, files...)
	}

	return result
}
