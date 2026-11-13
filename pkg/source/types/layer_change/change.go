package layer_change

type Change struct {
	Path        string
	Destination string
	Original    []byte
	Current     []byte
}
