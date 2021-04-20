package validation

import (
	"net/url"
	"regexp"
)

var snakeCase = regexp.MustCompile("^([a-z]*\\.)*[a-z]+(_[a-z]+)*$")

func IsValidUrl(addr string) bool {
	uri, err := url.Parse(addr)
	if err != nil {
		return false
	}
	return uri.Scheme != "" && uri.Host != ""
}

func IsValidKey(key string) bool {
	return snakeCase.MatchString(key)
}
