package scope

import (
	"phone-list/pkg/phone"

	"gorm.io/gorm"
)

var (
	getValidCodeRegex  = phone.GetValidCodeRegex
	getValidPhoneRegex = phone.GetValidPhoneRegex
)

func WherePhone(country string, state int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// if state = 1 the regEx will contain country filter
		if country != "" && state != phone.StateValid {
			db.Where("phone REGEXP ?", getValidCodeRegex(country))
		}

		if state == phone.StateValid {
			db.Where("phone REGEXP ?", getValidPhoneRegex(country))
		}

		if state == phone.StateInvalid {
			db.Where("phone NOT REGEXP ?", getValidPhoneRegex(country))
		}

		return db
	}
}
