package model

func mockGetValidPhoneRegex(regex string) {
	getValidPhoneRegex = func(country string) string {
		return regex
	}
}
