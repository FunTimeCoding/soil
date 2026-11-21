package probe

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"

func metapackageInstalled(lines []string) *bool {
	for _, line := range lines {
		if m := constant.InstalledPattern.FindStringSubmatch(line); m != nil {
			result := m[1] != constant.NotInstalled

			return &result
		}
	}

	return nil
}
