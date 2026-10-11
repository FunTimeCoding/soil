package packager

func WithOutputFile(v string) Option {
	return func(p *Packager) {
		p.OutputFile = v
	}
}
