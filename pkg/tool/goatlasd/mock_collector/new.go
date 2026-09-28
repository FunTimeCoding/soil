package mock_collector

func New(source string) *Collector {
	return &Collector{source: source}
}
