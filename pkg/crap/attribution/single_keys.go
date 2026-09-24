package attribution

import "sort"

func (m *Matrix) SingleKeys() []string {
	single := m.Single()
	result := make([]string, 0, len(single))

	for k := range single {
		result = append(result, k)
	}

	sort.Strings(result)

	return result
}
