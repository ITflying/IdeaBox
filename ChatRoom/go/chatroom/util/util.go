package util

import "regexp"

var (
	checkTel = regexp.MustCompile("^[0-9]*$")
	checkEmail = regexp.MustCompile("^@*$")
)

func IsTel(tel string) bool {
	return checkTel.MatchString(tel)
}

func IsEmail(email string) bool {
	return checkEmail.MatchString(email)
}