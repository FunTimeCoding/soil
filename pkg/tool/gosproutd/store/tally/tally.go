package tally

type Tally struct {
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
