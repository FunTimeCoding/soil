package result

type Group struct {
	Shape     string      `json:"shape"`
	Exemplar  string      `json:"exemplar"`
	Locations []*Location `json:"locations"`
}
