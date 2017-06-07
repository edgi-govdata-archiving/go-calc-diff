package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// fetchResult represents the result of a fetch
type fetchResult struct {
	url string
	res *http.Response
	err error
}

// FetchUrlString grabs urls a & b, returning their bodies as strings
func FetchUrlsString(a, b string) (aRes string, bRes string, err error) {
	ar, br, err := FetchUrls(a, b)
	if err != nil {
		return
	}

	defer ar.Body.Close()
	resBytes, err := ioutil.ReadAll(ar.Body)
	if err != nil {
		return
	}
	aRes = string(resBytes)

	defer br.Body.Close()
	resBytes, err = ioutil.ReadAll(ar.Body)
	if err != nil {
		return
	}
	bRes = string(resBytes)

	return
}

// FetchUrls grabs a & b in parallel
func FetchUrls(a, b string) (aRes *http.Response, bRes *http.Response, err error) {
	// create a channel to send responses on
	ch := make(chan fetchResult, 0)

	// fire off fetchurl in go routines
	go FetchUrl(a, ch)
	go FetchUrl(b, ch)

	// loop to block & read from parallel requests
	for i := 2; i > 0; i-- {
		// read from responses channel
		// this blocks until something sends on ch
		r := <-ch
		if r.err != nil {
			err = r.err
			return
		} else if r.url == a {
			aRes = r.res
		} else if r.url == b {
			bRes = r.res
		}
	}

	return
}

// FetchUrl grabs a url, returning the response on a chan of fetchResult
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
	}
	r.res = res
	// send the response to the done chan
	done <- r
}
