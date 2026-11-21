package security_tester

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_item"

func (o *Tester) Subjects(items []*remediation_item.Item) []string {
	o.t.Helper()
	result := make([]string, 0, len(items))

	for _, i := range items {
		result = append(result, i.Subject)
	}

	return result
}
