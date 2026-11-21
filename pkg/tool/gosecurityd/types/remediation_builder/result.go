package remediation_builder

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_count"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_result"
	"slices"
)

func (b *Builder) Result(flags []flag.Flag) *remediation_result.Result {
	decisions := make(map[string]*flag.Flag, len(flags))

	for j := range flags {
		decisions[flags[j].Subject] = &flags[j]
	}

	result := remediation_result.New()

	for _, i := range b.order {
		slices.Sort(i.Vulnerabilities)

		if i.Kind == constant.RemediationVulnerability {
			i.Count = affected_count.New()

			for _, host := range i.Hosts {
				i.Count.Add(b.worst[[2]string{i.Subject, host}].Kind())
			}
		}

		if d, okay := decisions[i.Subject]; okay {
			i.Decision = d
			i.Addressed = d.Mitigated
		}

		if i.Addressed {
			result.Addressed++
		} else {
			result.Open++
			result.Kinds[i.Kind]++
		}

		result.Items = append(result.Items, i)
	}

	result.Unimportant = len(b.unimportant)

	return result
}
