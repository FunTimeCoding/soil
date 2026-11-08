package census

func merge(
	into *reading,
	from *reading,
) {
	for k := range from.calls {
		into.calls[k] = true
	}

	for k := range from.fields {
		into.fields[k] = true
	}

	for k := range from.keys {
		into.keys[k] = true
	}

	for k := range from.packages {
		into.packages[k] = true
	}

	for l, methods := range from.uses {
		if into.uses[l] == nil {
			into.uses[l] = map[string]bool{}
		}

		for m := range methods {
			into.uses[l][m] = true
		}
	}

	into.atomically = into.atomically || from.atomically
}
