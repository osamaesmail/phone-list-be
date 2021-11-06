package service

import (
	"phone-list/app/repository"
	"phone-list/app/request"
	"phone-list/app/response"
	"phone-list/pkg/logger"
)

type CustomerService interface {
	List(req request.CustomerListRequest) (*response.CustomerListResponse, error)
}

func NewCustomerService(customerRepository repository.CustomerRepository) CustomerService {
	return &customerService{customerRepository}
}

type customerService struct {
	customerRepository repository.CustomerRepository
}

func (s *customerService) List(req request.CustomerListRequest) (*response.CustomerListResponse, error) {
	customers, err := s.customerRepository.List(req.Limit, req.Offset, req.Country, req.State)
	if err != nil {
		logger.Log().Err(err).Msg("Failed to get customers list")
		return nil, err
	}

	total, err := s.customerRepository.Count(req.Country, req.State)
	if err != nil {
		logger.Log().Err(err).Msg("Failed to get customers count")
		return nil, err
	}

	return response.NewCustomerListResponse(customers, total, req.Limit, req.Offset), nil
}
