package snapshot

import "time"

type Snapshot struct {
	Instance       string    `gorm:"primaryKey;column:instance"`
	Project        int64     `gorm:"primaryKey;autoIncrement:false;column:project"`
	Request        int64     `gorm:"primaryKey;autoIncrement:false;column:request"`
	Update         time.Time `gorm:"column:request_updated_at"`
	State          string    `gorm:"column:state"`
	Draft          bool      `gorm:"column:draft"`
	MergeStatus    string    `gorm:"column:merge_status"`
	Pipeline       int64     `gorm:"column:pipeline"`
	PipelineStatus string    `gorm:"column:pipeline_status"`
	Commit         string    `gorm:"column:commit_sha"`
	Notes          int64     `gorm:"column:notes"`
	LastNote       int64     `gorm:"column:last_note"`
	Approvers      string    `gorm:"column:approvers"`
}
