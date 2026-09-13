package user

type User struct {
	Account           string `json:"account"`
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	Mail              string `json:"mail"`
	Unique            string `json:"unique"`
	DistinguishedName string `json:"distinguished_name"`
}
