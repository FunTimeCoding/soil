package response

type FormulaService struct {
	Name                 ServiceName       `json:"name,omitempty"`
	Run                  any               `json:"run,omitempty"`
	RunType              string            `json:"run_type,omitempty"`
	KeepAlive            KeepAlive         `json:"keep_alive,omitempty"`
	EnvironmentVariables map[string]string `json:"environment_variables,omitempty"`
	WorkingDirectory     string            `json:"working_dir,omitempty"`
	LogPath              string            `json:"log_path,omitempty"`
	ErrorLogPath         string            `json:"error_log_path,omitempty"`
	RequireRoot          bool              `json:"require_root,omitzero"`
}
