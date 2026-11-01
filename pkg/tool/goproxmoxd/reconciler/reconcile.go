package reconciler

func (r *Reconciler) Reconcile() {
	for _, hypervisor := range r.service.Inventory().Names() {
		p, e := r.service.Reconcile(r.context, hypervisor, false, true)

		if e != nil {
			r.fail(hypervisor, e)

			continue
		}

		for _, c := range p.Changes {
			r.logger.Plain(
				"reconcile %s: %s %s -> %s",
				hypervisor,
				c.Name,
				c.From,
				c.To,
			)
		}
	}
}
