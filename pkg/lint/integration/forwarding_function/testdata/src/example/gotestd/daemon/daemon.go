package daemon

import (
	"example/gotestd/guard"
	"example/gotestd/service"
	"example/gotestd/service/loan"
)

type Daemon struct {
	service *service.Service
	loan    *loan.Service
	guard   *guard.Guard
}

func (d *Daemon) Packages(limit int) []string {
	return d.service.Packages(limit)
}

func (d *Daemon) Renew(identifier int) error {
	return d.loan.Renew(identifier)
}

func (d *Daemon) Wait(timeout int) error {
	return d.guard.Wait(timeout)
}
