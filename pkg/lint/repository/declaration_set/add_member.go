package declaration_set

func (s *Set) addMember(
	owner string,
	member string,
) {
	if s.members[owner] == nil {
		s.members[owner] = map[string]bool{}
	}

	s.members[owner][member] = true
}
