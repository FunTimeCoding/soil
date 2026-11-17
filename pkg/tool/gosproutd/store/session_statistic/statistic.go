package session_statistic

type Statistic struct {
	Session    string
	Emitted    int
	Answered   int
	Bypassed   int
	Declined   int
	Irrelevant int
	Postponed  int
	Defaulted  int
	Abandoned  int
	Bumped     int
	Turns      int
}
