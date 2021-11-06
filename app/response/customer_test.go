package response

import (
	"testing"

	"phone-list/app/model"

	"github.com/stretchr/testify/assert"
)

func TestNewCustomerResponse(t *testing.T) {
	c := &model.Customer{
		ID:    1,
		Name:  "foo",
		Phone: "bar",
	}

	resp := NewCustomerResponse(c)

	assert.Equal(t, uint(1), resp.ID)
	assert.Equal(t, "bar", resp.Phone)
	assert.Equal(t, false, resp.PhoneValid)
}

func TestNewCustomerListResponse(t *testing.T) {
	c1 := &model.Customer{
		ID:    1,
		Name:  "foo",
		Phone: "bar",
	}

	c2 := &model.Customer{
		ID:    1,
		Name:  "foo",
		Phone: "bar",
	}

	payloads := []*model.Customer{c1, c2}

	resp := NewCustomerListResponse(payloads, 3, 2, 1)

	assert.Equal(t, 3, resp.Total)
	assert.Equal(t, 2, resp.Limit)
	assert.Equal(t, 1, resp.Offset)
	assert.Equal(t, c1.ID, resp.Rows[0].ID)

	assert.Len(t, resp.Rows, 2)

	assert.Equal(t, c1.Phone, resp.Rows[0].Phone)
	assert.Equal(t, c2.ID, resp.Rows[1].ID)
	assert.Equal(t, c2.Phone, resp.Rows[1].Phone)
}
