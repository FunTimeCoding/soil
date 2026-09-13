package argument

type CreateUser struct {
	Account  string `json:"account"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}
