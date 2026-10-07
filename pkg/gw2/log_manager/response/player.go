package response

type Player struct {
	Name                string `json:"Name"`
	AccountName         string `json:"AccountName"`
	Subgroup            int    `json:"Subgroup"`
	Profession          int    `json:"Profession"`
	EliteSpecialization int    `json:"EliteSpecialization"`
	GuildGUID           string `json:"GuildGuid"`
	Tag                 int    `json:"Tag"`
}
