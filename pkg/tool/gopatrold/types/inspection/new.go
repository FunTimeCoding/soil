package inspection

func New(host string) *Inspection {
	return &Inspection{Host: host}
}
