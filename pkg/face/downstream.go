package face

type Downstream interface {
	Trigger(changes []string) error
}
