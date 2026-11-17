package service_tester

func WithEmbedModel(name string) Option {
	return func(o *Tester) {
		o.EmbedModel = name
	}
}
