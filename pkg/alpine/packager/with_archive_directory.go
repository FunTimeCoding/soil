package packager

func WithArchiveDirectory(v string) Option {
	return func(p *Packager) {
		p.ArchiveDirectory = v
	}
}
