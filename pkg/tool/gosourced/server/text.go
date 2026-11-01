package server

func text(p *string) string {
	if p == nil {
		return ""
	}

	return *p
}
