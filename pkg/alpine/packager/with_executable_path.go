package packager

func WithExecutablePath(v string) Option {
	return func(p *Packager) {
		p.ExecutablePath = v
	}
}
