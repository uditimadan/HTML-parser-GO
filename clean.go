package main

import (
	"net/url"
)

//cleans a single href based on the given host
func clean(host string, href string) string {
	baseURL, err := url.Parse(host)
	if err != nil {
		return ""
	}
	u, err := url.Parse(href)
	if err != nil {
		return href
	}
	return baseURL.ResolveReference(u).String() //returns the fully resolved URL
}
