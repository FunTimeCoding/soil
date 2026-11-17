package argument

type CreateAddress struct {
	Device    string `json:"device"`
	Interface string `json:"interface"`
	Address   string `json:"address"`
	Status    string `json:"status"`
}
