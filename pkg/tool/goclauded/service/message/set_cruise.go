package message

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/refusal"
	"time"
)

func (s *Service) SetCruise(
	sessionIdentifier string,
	mode constant.Cruise,
	pace time.Duration,
) error {
	switch mode {
	case constant.CruiseOff, constant.CruiseOn:
		pace = 0
	case constant.CruisePaced:
		if pace%time.Minute != 0 ||
			pace < constant.CruisePaceMinimumMinutes*time.Minute ||
			pace > constant.CruisePaceMaximumMinutes*time.Minute {
			return refusal.New(
				fmt.Sprintf(
					"a paced cruise needs a whole number of minutes from %d to %d, not %g",
					constant.CruisePaceMinimumMinutes,
					constant.CruisePaceMaximumMinutes,
					pace.Minutes(),
				),
			)
		}
	default:
		return refusal.New(
			fmt.Sprintf("cruise mode must be off, on or paced, not %s", mode),
		)
	}

	if e := s.store.SetCruise(sessionIdentifier, mode, pace); e != nil {
		return e
	}

	s.notify()

	return nil
}
