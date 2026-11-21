package remediation_builder

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"slices"
)

func (b *Builder) Add(
	key string,
	s *system_result.Summary,
	f *system_result.Finding,
) {
	i := b.Item(key, s, f)

	if pair := [2]string{key, f.Vulnerability}; !b.found[pair] {
		b.found[pair] = true
		i.Vulnerabilities = append(i.Vulnerabilities, f.Vulnerability)
	}

	i.TopScore = max(i.TopScore, f.Score)

	if i.Kind != constant.RemediationVulnerability {
		return
	}

	pair := [2]string{key, s.Name}

	if current := b.worst[pair]; current == nil ||
		slices.Index(constant.KindOrder, f.Kind()) >
			slices.Index(constant.KindOrder, current.Kind()) {
		b.worst[pair] = f
	}
}
