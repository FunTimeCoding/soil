package option

func New() *Postgres {
	return &Postgres{Client: BuildClient}
}
