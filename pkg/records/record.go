package records

import (
	"fmt"

	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	UUID     string `gorm:"uniqueIndex" json:"uuid"`
	Domain   string `gorm:"index" json:"-"`
	Category string `gorm:"index" json:"-"`
	Body     string `json:"body"`
}

func (r Record) Validate() error {
	if r.UUID == "" {
		return fmt.Errorf("UUID is required")
	}
	if r.Body == "" {
		return fmt.Errorf("Body is required")
	}
	return nil
}
