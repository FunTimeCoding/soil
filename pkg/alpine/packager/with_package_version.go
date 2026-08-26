package packager

func WithPackageVersion(v string) Option {
	return func(p *Packager) {
		p.PackageVersion = v
	}
}
