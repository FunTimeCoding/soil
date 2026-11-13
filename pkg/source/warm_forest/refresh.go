package warm_forest

func (f *Forest) Refresh() map[string][]string {
	result := make(map[string][]string, len(f.trees))

	for root, t := range f.trees {
		result[root] = t.Refresh()
	}

	return result
}
