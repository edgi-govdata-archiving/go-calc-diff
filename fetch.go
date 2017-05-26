package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

// FetchUrl grabs a url, returning it's response body
func FetchUrl(rawurl string) ([]byte, error) {
	// if the json doc specifies a url value, grab it from the internet
	parsed, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	if !(parsed.Scheme == "http" || parsed.Scheme == "https") {
		return nil, nil
	}

	res, err := http.Get(parsed.String())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
