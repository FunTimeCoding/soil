package packager

func WithControlDirectory(v string) Option {
	return func(p *Packager) {
		p.ControlDirectory = v
	}
}
