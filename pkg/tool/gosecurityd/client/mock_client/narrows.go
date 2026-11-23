package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"

func narrows(s *system_selection.Selection) bool {
	return s.Show != "" || len(s.Labels) > 0 || s.Search != ""
}
