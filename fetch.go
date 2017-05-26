package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// fetchResult represents the result of a fetch
type fetchResult struct {
	url  string
	body string
	err  error
}

// FetchUrls grabs a & b in parallel
func FetchUrls(a, b string) (aRes string, bRes string, err error) {
	// create a channel to send responses on
	ch := make(chan fetchResult, 0)

	// fire off fetchurl in go routines
	go FetchUrl(a, ch)
	go FetchUrl(b, ch)

	// loop
	for i := 2; i > 0; i-- {
		// read from responses channel
		// this blocks until something sends on ch
		res := <-ch
		if res.err != nil {
			err = res.err
			return
		} else if res.url == a {
			aRes = res.body
		} else if res.url == b {
			bRes = res.body
		}
	}

	return
}

// FetchUrl grabs a url, returning it's response body
func FetchUrl(rawurl string, done chan fetchResult) {
	// our result, labelled by the rawurl
	r := fetchResult{url: rawurl}

	// if the json doc specifies a url value, grab it from the internet
	parsed, err := url.Parse(rawurl)
	if err != nil {
		r.err = fmt.Errorf("invalid url: %s", rawurl)
		done <- r
		return
	}
	if !(parsed.Scheme == "http" || parsed.Scheme == "https") {
		r.err = fmt.Errorf("differ can only fetch http or https urls")
		done <- r
		return
	}

	// GET the passed in url
	res, err := http.Get(parsed.String())
	if err != nil {
		r.err = fmt.Errorf("fetch url %s failed: %s", rawurl, err)
		done <- r
		return
	}
	defer res.Body.Close()

	// read the
	resBytes, err := ioutil.ReadAll(res.Body)
	r.body = string(resBytes)
	r.err = err
	// send the response to the done chan
	done <- r
}
