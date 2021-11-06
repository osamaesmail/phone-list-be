package response

import (
	"phone-list/app/model"
)

type CustomerResponse struct {
	ID         uint   `json:"id"`
	Phone      string `json:"phone"`
	PhoneValid bool   `json:"phone_valid"`
}

type CustomerListResponse struct {
	paginationResponse
	Rows []*CustomerResponse `json:"rows"`
}

func NewCustomerResponse(payload *model.Customer) *CustomerResponse {
	res := &CustomerResponse{
		ID:         payload.ID,
		Phone:      payload.Phone,
		PhoneValid: payload.IsValidPhone(),
	}
	return res
}

func NewCustomerListResponse(payloads []*model.Customer, total, limit, offset int) *CustomerListResponse {
	rows := make([]*CustomerResponse, len(payloads))
	for i, payload := range payloads {
		rows[i] = NewCustomerResponse(payload)
	}
	res := &CustomerListResponse{
		paginationResponse: paginationResponse{
			Limit:  limit,
			Offset: offset,
			Total:  total,
		},
		Rows: rows,
	}

	return res
}
