package utils

import (
	url2 "net/url"
)

func GetUrlFromString(urlStr string) url2.URL {
	url, err := url2.Parse(urlStr)
	CheckError(err)

	return *url
}
