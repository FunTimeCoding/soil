package service_tester

import "github.com/funtimecoding/soil/pkg/chat/teams/user"

func (o *Tester) SeedPeople() {
	o.Source.SetMe(user.New("me-1", "Alfa", "alfa@example.com"))
	o.Source.AddUser(user.New("user-9", "Juliett", "juliett@example.com"))
	o.Source.AddUser(
		user.New("user-8", "Juliett Romeo", "juliett.romeo@example.com"),
	)
	o.Source.AddUser(user.New("user-7", "Kilo", "kilo@example.com"))
}
