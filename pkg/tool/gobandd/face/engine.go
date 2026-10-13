package face

import "github.com/funtimecoding/soil/pkg/band/response"

type Engine interface {
	GeneralSettings() (*response.GeneralSettings, error)
	Setup() (*response.Setup, error)
	ScreenSettings() (*response.ScreenSettings, error)
	Redirection() (*response.Redirection, error)
	ScreenService() (*response.ScreenService, error)
	Consent() (*response.Consent, error)
	PowerState() (*response.PowerState, error)
	RequestPowerState(state int) error
	EnableListener() error
	EnableScreenService() error
}
