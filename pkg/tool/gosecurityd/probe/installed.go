package probe

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"strings"
)

func installed(
	lines []string,
	upgradable map[string]string,
) []*client.InstalledPackage {
	var result []*client.InstalledPackage

	for _, line := range lines {
		f := strings.Split(line, "\t")

		if len(f) != constant.PackageFields ||
			len(f[0]) < 2 ||
			f[0][1] != constant.InstalledStatus {
			continue
		}

		p := &client.InstalledPackage{
			Name:          f[1],
			Version:       f[2],
			Source:        f[3],
			SourceVersion: f[4],
		}

		if c, okay := upgradable[f[1]]; okay {
			p.Candidate = &c
		}

		result = append(result, p)
	}

	return result
}
