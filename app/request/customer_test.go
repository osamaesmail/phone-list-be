package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type contextMock struct{}

func (*contextMock) BindQuery(obj interface{}) error {
	req := obj.(*CustomerListRequest)
	req.Country = "foo"
	req.State = 1
	req.Limit = 2
	req.Offset = 1

	return nil
}

func TestNewCustomerListRequest(t *testing.T) {
	c := &contextMock{}

	req, _ := NewCustomerListRequest(c)

	assert.Equal(t, "foo", req.Country)
	assert.Equal(t, 1, req.State)
	assert.Equal(t, 2, req.Limit)
	assert.Equal(t, 1, req.Offset)
}
