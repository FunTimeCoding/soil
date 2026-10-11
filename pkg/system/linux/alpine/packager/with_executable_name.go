package packager

func WithExecutableName(v string) Option {
	return func(p *Packager) {
		p.ExecutableName = v
	}
}
