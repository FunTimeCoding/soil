package remediation_builder

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/subject"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_item"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
)

func (b *Builder) Item(
	key string,
	s *system_result.Summary,
	f *system_result.Finding,
) *remediation_item.Item {
	i, okay := b.items[key]

	if !okay {
		k := subject.Kind(key)
		release := s.Release
		source := ""

		switch k {
		case constant.RemediationVulnerability:
			release = ""
		case constant.RemediationUpgrade:
			source = f.Label()
		}

		i = remediation_item.New(key, k, release, source)
		i.Kernel = k == constant.RemediationUpgrade && f.Kernel

		if k == constant.RemediationKernel {
			i.Reason = s.KernelReason
		}

		b.items[key] = i
		b.order = append(b.order, i)
	}

	if pair := [2]string{key, s.Name}; !b.hosts[pair] {
		b.hosts[pair] = true
		i.Hosts = append(i.Hosts, s.Name)
	}

	return i
}
