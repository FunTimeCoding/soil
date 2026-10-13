package client

import (
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"time"
)

func New(key string) *Client {
	p := stall_client.New().Timeout(10 * time.Second).Build()
	l := Configuration(key, p.Client())

	return &Client{
		pool:         p,
		Account:      Account(l),
		Alert:        Alert(l),
		Contact:      Contact(l),
		Deployment:   Deployment(l),
		Escalation:   Escalation(l),
		Forward:      Forward(l),
		Heartbeat:    Heartbeat(l),
		Incident:     Incident(l),
		Integration:  Integration(l),
		Log:          Log(l),
		Maintenance:  Maintenance(l),
		Notification: Notification(l),
		Policy:       Policy(l),
		Schedule:     Schedule(l),
		Service:      Service(l),
		Team:         Team(l),
		User:         User(l),
	}
}
