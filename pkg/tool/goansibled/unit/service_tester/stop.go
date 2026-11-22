package service_tester

func (o *Tester) Stop() {
	o.Service()
	o.runner.Close()
	<-o.runner.Finished()
}
