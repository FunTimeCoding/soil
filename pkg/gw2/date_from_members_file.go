package gw2

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func dateFromMembersFile(file string) string {
	//  Example: ~/Downloads/members_Theres_Always_A_Bigger_Fish_BAIT_2024-11-01_13-21-05.json
	last := strings.LastIndex(file, constant.Underscore)
	secondLast := strings.LastIndex(file[:last], constant.Underscore)

	return file[secondLast+1 : last]
}
