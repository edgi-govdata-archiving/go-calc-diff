package main

import (
	"testing"
)

func TestFetchUrl(t *testing.T) {
	cases := []struct {
		url string
		err error
	}{
		{"ftp://stuff.ftp.dontfetch.com", ErrBadUrl},
		{"http://www.apple.com", nil},
		{"https://www.apple.com", nil},
	}

	for i, c := range cases {
		res := make(chan fetchResult, 0)
		go FetchUrl(c.url, res)
		r := <-res

		if r.err != c.err {
			t.Errorf("case %d error mismatch. expected: '%s', got: '%s'", i, c.err, r.err)
		}

		if r.res != nil {
			r.res.Body.Close()
		}
	}
}
