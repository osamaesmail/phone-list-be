package service

import (
	"phone-list/app/model"
	"phone-list/app/repository"
	"phone-list/app/request"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCustomerService(t *testing.T) {
	repo := &repository.CustomerRepositoryMock{}
	service := NewCustomerService(repo).(*customerService)

	assert.Equal(t, service.customerRepository, repo)
}

func TestCustomerService_List(t *testing.T) {
	c := &model.Customer{
		ID:    1,
		Name:  "foo",
		Phone: "bar",
	}
	repo := &repository.CustomerRepositoryMock{
		Total:    3,
		Payloads: []*model.Customer{c, c},
	}
	service := NewCustomerService(repo)

	req := request.CustomerListRequest{}

	resp, _ := service.List(req)

	assert.Equal(t, 2, len(resp.Rows))
	assert.Equal(t, 3, resp.Total)
}
