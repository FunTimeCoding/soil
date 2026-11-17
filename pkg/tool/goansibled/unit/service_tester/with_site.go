package service_tester

func WithSite(site string) Option {
	return func(o *Tester) {
		o.Option.Site = site
	}
}
