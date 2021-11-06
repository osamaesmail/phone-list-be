package scope

func mockGetValidCodeRegex(regex string) {
	getValidCodeRegex = func(country string) string {
		return regex
	}
}

func mockGetValidPhoneRegex(regex string) {
	getValidPhoneRegex = func(country string) string {
		return regex
	}
}
