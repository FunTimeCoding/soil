package argument

type SetPassword struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
