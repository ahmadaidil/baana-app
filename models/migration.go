package models

// Migration .
type Migration struct {
	Version string `gorm:"type:varchar(255);primary_key" json:"version" form:"version"`
}
