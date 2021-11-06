package phone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValidPhoneRegex(t *testing.T) {
	assert.Equal(t, "\\(237\\)\\ ?[2368]\\d{7,8}$", GetValidPhoneRegex("cameroon"))

	assert.Contains(t, GetValidPhoneRegex(""), "(\\(237\\)\\ ?[2368]\\d{7,8}$)|")
}

func TestGetValidCodeRegex(t *testing.T) {
	assert.Equal(t, "\\(237\\)\\.*", GetValidCodeRegex("cameroon"))

	assert.Equal(t, "", GetValidCodeRegex(""))
}
