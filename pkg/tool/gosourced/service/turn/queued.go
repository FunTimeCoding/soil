package turn

func (s *Scheduler) queued() int {
	result := 0

	for _, queue := range s.waiting {
		result += len(queue)
	}

	return result
}
