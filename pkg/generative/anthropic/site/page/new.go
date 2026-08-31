package page

func New(
	sessionPercent int,
	sessionReset string,
	weeklyAllPercent int,
	weeklyAllReset string,
	fablePercent int,
	fableReset string,
) *Usage {
	return &Usage{
		SessionPercent:   sessionPercent,
		SessionReset:     sessionReset,
		WeeklyAllPercent: weeklyAllPercent,
		WeeklyAllReset:   weeklyAllReset,
		FablePercent:     fablePercent,
		FableReset:       fableReset,
	}
}
