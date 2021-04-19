package main

import (
	"net/url"
	"regexp"
)

var snakeCaseRE = regexp.MustCompile("^[a-z]+(_[a-z]+)*$")

func isValidUrl(addr string) bool {
	uri, err := url.Parse(addr)
	if err != nil {
		return false
	}
	return uri.Scheme != "" && uri.Host != ""
}

func isValidKey(key string) bool {
	return snakeCaseRE.MatchString(key)
}