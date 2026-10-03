package strings

func Compare(
	past []string,
	now []string,
) ([]string, []string, []string) {
	add := filterMembership(now, past, false)
	remove := filterMembership(past, now, false)
	stay := filterMembership(past, now, true)

	return add, remove, stay
}
