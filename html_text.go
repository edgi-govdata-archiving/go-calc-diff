package main

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

// HtmlTextContent is a very naive text-ripper. It constructs a jQuery-like document
// from returned html, grabs all elements from the body, removes all "script-like" tags
// from the selection, and concatenates the text values into a string with space-separators.
// surprisingly effective.
func HtmlTextContent(res *http.Response) (string, error) {
	// response comes from remote server, must be closed when we're done
	defer res.Body.Close()
	// goquery is the package that does most of the heavy lifting here
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return "", err
	}

	// slice of text elements
	text := []string{}

	// iterate through designated body nodes, adding any text that isn't blank to the text slice
	doc.Find("body").Find("*").Not("script, style, noscript, canvas").Each(func(i int, el *goquery.Selection) {
		t := strings.TrimSpace(el.Text())
		if t != "" {
			text = append(text, t)
		}
	})

	return strings.Join(text, " "), nil
}
