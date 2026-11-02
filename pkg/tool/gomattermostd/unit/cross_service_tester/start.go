package cross_service_tester

func (o *Tester) Start() {
	for _, w := range o.Daemon.Workers() {
		w.Start()
		o.t.Cleanup(w.Close)
	}
}
