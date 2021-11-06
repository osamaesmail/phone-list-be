package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomer_IsValidPhoneTrue(t *testing.T) {
	mockGetValidPhoneRegex("foo")

	c := Customer{
		Phone: "foo",
	}

	assert.True(t, c.IsValidPhone(), "IsValidPhone")
}

func TestCustomer_IsValidPhoneFalse(t *testing.T) {
	mockGetValidPhoneRegex("foo")

	c := Customer{
		Phone: "bar",
	}

	assert.False(t, c.IsValidPhone())
}
