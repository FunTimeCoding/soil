package argument

type ModifyUser struct {
	Account string `json:"account"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Mail    string `json:"mail"`
}
