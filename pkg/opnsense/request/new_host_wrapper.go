package request

func NewHostWrapper(h *Host) *HostWrapper {
	return &HostWrapper{Host: h}
}
