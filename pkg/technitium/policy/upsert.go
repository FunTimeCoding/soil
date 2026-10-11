package policy

func Upsert(
	v []*Policy,
	p *Policy,
) []*Policy {
	for i, e := range v {
		if e.KeyName == p.KeyName && e.Domain == p.Domain {
			v[i] = p

			return v
		}
	}

	return append(v, p)
}
