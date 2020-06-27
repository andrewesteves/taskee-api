package utils

import "regexp"

// Contains checks if a value exists inside an array
func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// CheckEmail to verify if is a valid email
func CheckEmail(email string) bool {
	pattern := "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	isEmail := regexp.MustCompile(pattern)
	return isEmail.MatchString(email)
}
