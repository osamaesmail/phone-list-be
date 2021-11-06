package repository

import (
	"gorm.io/gorm"
)

func mockScopeWhere() {
	scopeWherePhone = func(country string, state int) func(db *gorm.DB) *gorm.DB {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("country = ?", country).Where("state = ?", state)
		}
	}
}
