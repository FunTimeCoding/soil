package not_configured

func (e *NotConfiguredError) Error() string { return e.Message }
