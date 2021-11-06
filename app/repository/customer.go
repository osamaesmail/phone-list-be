package repository

import (
	"phone-list/app/model"
	"phone-list/app/scope"

	"gorm.io/gorm"
)

var scopeWherePhone = scope.WherePhone

type CustomerRepository interface {
	List(limit, offset int, country string, state int) ([]*model.Customer, error)
	Count(country string, state int) (int, error)
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

type customerRepository struct {
	db *gorm.DB
}

func (repo *customerRepository) List(limit, offset int, country string, state int) ([]*model.Customer, error) {
	var customers []*model.Customer

	err := repo.db.
		Scopes(scopeWherePhone(country, state)).
		Offset(offset).Limit(limit).Find(&customers).Error
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (repo *customerRepository) Count(country string, state int) (int, error) {
	var count int64

	err := repo.db.Model(&model.Customer{}).
		Scopes(scopeWherePhone(country, state)).
		Count(&count).Error
	if err != nil {
		return int(count), err
	}

	return int(count), nil
}
