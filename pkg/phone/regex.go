package phone

import (
	"fmt"
	"strings"
)

const (
	phoneRegexStr = "\\(%d\\)\\ ?%s$"
	codeRegexStr  = "\\(%d\\)\\.*"
)

type countryRegex struct {
	Code         int
	Name         string
	LeadingRegex string
}

func newCountryRegex(code int, name, leadingRegex string) *countryRegex {
	return &countryRegex{
		Code:         code,
		Name:         name,
		LeadingRegex: leadingRegex,
	}
}

var countryRegexMap = map[string]*countryRegex{
	"cameroon":   newCountryRegex(237, "cameroon", "[2368]\\d{7,8}"),
	"ethiopia":   newCountryRegex(251, "ethiopia", "[1-59]\\d{8}"),
	"morocco":    newCountryRegex(212, "morocco", "[5-9]\\d{8}"),
	"mozambique": newCountryRegex(258, "mozambique", "[28]\\d{7,8}"),
	"uganda":     newCountryRegex(256, "uganda", "\\d{9}"),
}

func GetValidPhoneRegex(country string) string {
	if countryRegex, exists := countryRegexMap[country]; exists {
		return fmt.Sprintf(phoneRegexStr, countryRegex.Code, countryRegex.LeadingRegex)
	}

	regExList := make([]string, len(countryRegexMap))
	listIndex := 0
	for c := range countryRegexMap {
		regExList[listIndex] = GetValidPhoneRegex(c)
		listIndex++
	}

	return fmt.Sprintf("(%s)", strings.Join(regExList, ")|("))
}

func GetValidCodeRegex(country string) string {
	if countryRegex, exists := countryRegexMap[country]; exists {
		return fmt.Sprintf(codeRegexStr, countryRegex.Code)
	}

	return "x\\by"
}
