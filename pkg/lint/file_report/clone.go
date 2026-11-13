package file_report

func (r *Report) Clone() *Report {
	result := &Report{
		FilePath: r.FilePath,
		Original: r.Original,
		Fixed:    r.Fixed,
		modified: r.modified,
	}

	for _, c := range r.Concerns {
		result.Concerns = append(result.Concerns, c.Clone())
	}

	return result
}
