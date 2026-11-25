package snapshot

import "strings"

func (s *Snapshot) ApproverList() []string {
	if s.Approvers == "" {
		return nil
	}

	return strings.Split(s.Approvers, ",")
}
