package model

import (
	"phone-list/pkg/phone"
	"regexp"
)

var getValidPhoneRegex = phone.GetValidPhoneRegex

type Customer struct {
	ID uint `gorm:"primarykey"`

	Name  string
	Phone string
}

func (*Customer) TableName() string {
	return "customer"
}

func (c *Customer) IsValidPhone() bool {
	if r, err := regexp.Compile(getValidPhoneRegex("")); err == nil {
		return r.MatchString(c.Phone)
	}

	return false
}
