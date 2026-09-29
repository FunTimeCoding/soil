package job

func (j *Job) ProjectIdentifier() int64 {
	if j.Raw.Project != nil && j.Raw.Project.ID != 0 {
		return j.Raw.Project.ID
	}

	return j.Raw.Pipeline.ProjectID
}
