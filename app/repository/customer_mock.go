package repository

import "phone-list/app/model"

type CustomerRepositoryMock struct {
	CustomerRepository
	Total    int
	Payloads []*model.Customer
}

func (repo *CustomerRepositoryMock) List(limit, offset int, country string, state int) ([]*model.Customer, error) {
	return repo.Payloads, nil
}

func (repo *CustomerRepositoryMock) Count(country string, state int) (int, error) {
	return repo.Total, nil
}
