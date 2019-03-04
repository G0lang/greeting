package app

import "github.com/jinzhu/gorm"

// Name Model
type Name struct {
	gorm.Model
	Name string `gorm:"column:name" json:"name"`
}

// TableName table Name on db
func (b *Name) TableName() string {
	return "name"
}

// InsertName to db
func InsertName(db *gorm.DB, n *Name) (err error) {
	if err = db.Save(n).Error; err != nil {
		return err
	}
	return nil
}
