package model

func New(reader Reader) *Model {
	return &Model{reader: reader}
}
