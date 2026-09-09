package option

type Monitor struct {
	Enabled             bool
	Schedule            string
	Channels            []string
	Topics              []string
	MessageLimit        int
	NotificationChannel string
	Username            string
}
