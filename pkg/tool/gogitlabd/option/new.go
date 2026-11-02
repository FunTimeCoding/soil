package option

func New() *Gitlab {
	return &Gitlab{Client: BuildClient}
}
