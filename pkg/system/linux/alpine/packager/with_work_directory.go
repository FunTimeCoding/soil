package packager

func WithWorkDirectory(v string) Option {
	return func(p *Packager) {
		p.WorkDirectory = v
	}
}
