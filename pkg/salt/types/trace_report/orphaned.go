package trace_report

func orphaned(nodes []*Node) int {
	var result int

	for _, n := range nodes {
		if n.Orphaned {
			result++
		}
	}

	return result
}
