package instance

func (i *Instance) InstanceHost() string {
	return i.Alertmanager.Host
}
