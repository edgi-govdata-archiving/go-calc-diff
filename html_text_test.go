package main

import (
	"bytes"
	"fmt"
	"github.com/qri-io/go-diff/diffmatchpatch"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestHtmlTextContent(t *testing.T) {
	cases := []struct {
		testHtml   string
		expectText string
		err        error
	}{
		{"simple.html", "simple.txt", nil},
		{"epa.gov.html", "epa.gov.txt", nil},
	}

	for i, c := range cases {
		in, err := ioutil.ReadFile(filepath.Join("test_files", c.testHtml))
		if err != nil {
			t.Errorf("case %d error reading test html file: %s", i, err.Error())
			continue
		}

		outbytes, err := ioutil.ReadFile(filepath.Join("test_files", c.expectText))
		if err != nil {
			t.Errorf("case %d error reading test output text file: %s", i, err.Error())
			continue
		}
		out := string(outbytes)

		// create a fake response from the file
		got, err := HtmlTextContent(&http.Response{
			Request: httptest.NewRequest("GET", "http://testhtmltextcontent.com", nil),
			Body:    ioutil.NopCloser(bytes.NewBuffer(in)),
		})
		if err != c.err {
			t.Errorf("case %d error mismatch: %s != %s", i, c.err, err)
		}

		// Uncomment this to write the raw text to output-[casenum].txt
		// ioutil.WriteFile(fmt.Sprintf("output-%d.txt", i), []byte(got), os.ModePerm)

		if got != out {
			// t.Errorf("case %d text output mismatch", i)

			filename := fmt.Sprintf("errors-%d.html", i)
			dmp := diffmatchpatch.New()
			ioutil.WriteFile(filename, []byte(dmp.DiffPrettyHtml(dmp.DiffMain(got, out, false))), os.ModePerm)
			t.Errorf("case %d text output mismatch diff of text written to: %s", i, filename)
		}
	}
}
