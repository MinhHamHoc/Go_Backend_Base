package provider

import (
	"regexp"
	"time"
)

var regexPhoneNumber = regexp.MustCompile(`(?m)(84|0[0-9])+([0-9]{8})\b`)

func IsPhoneNumber(str string) bool {
	return regexPhoneNumber.MatchString(str)
}
func TimeInUTC(t time.Time) time.Time {
	return t.UTC()
}

func TimeInLocal(t time.Time) time.Time {
	return t.In(time.Local)
}
