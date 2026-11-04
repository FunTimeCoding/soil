package reported

func (e *ReportedError) Error() string { return e.Message }
