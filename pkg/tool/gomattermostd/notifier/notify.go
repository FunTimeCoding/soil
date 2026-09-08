package notifier

func (n *Notifier) Notify(callsign string, body string) {
	if callsign == "" {
		return
	}

	if e := n.connector.Notify(callsign, n.source, body); e != nil {
		n.reporter.CaptureException(e)
	}
}
