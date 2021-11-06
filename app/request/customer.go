package request

import "phone-list/context"

type CustomerListRequest struct {
	paginationRequest
	State   int    `form:"state"`
	Country string `form:"country"`
}

func NewCustomerListRequest(c context.Context) (CustomerListRequest, error) {
	var req CustomerListRequest
	err := c.BindQuery(&req)
	if err != nil {
		return req, err
	}

	return req, nil
}
