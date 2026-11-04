package subscribed

func (o *Subscribed) Close() {
	o.Roster.Close()
}
