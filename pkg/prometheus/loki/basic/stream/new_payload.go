package stream

func NewPayload(streams ...*Stream) *Payload {
	return &Payload{Streams: streams}
}
