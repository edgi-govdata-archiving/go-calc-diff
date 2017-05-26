package main

import (
	"testing"
)

func TestFetchUrl(t *testing.T) {
	cases := []struct {
		url         string
		shouldFetch bool
		err         error
	}{
		{"ftp://stuff.ftp.dontfetch.com", false, nil},
		{"http://www.apple.com", true, nil},
		{"https://www.apple.com", true, nil},
	}

	for i, c := range cases {
		_, got := FetchUrl(c.url)

		if got != c.err {
			t.Errorf("case %d error mismatch. expected: '%s', got: '%s'", i, c.err, got)
		}

		// if fetched != c.shouldFetch {
		// 	t.Errorf("case %d error fetched mismatch. expected: %t, got: %t", i, c.shouldFetch, fetched)
		// }

	}
}
