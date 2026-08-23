package face

type Instrument interface {
	Recorder() Recorder
	Reporter() Reporter
}
