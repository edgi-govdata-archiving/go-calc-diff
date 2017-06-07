package main

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	// "strings"
)

// HtmlTextContent is a very naive text-ripper. It constructs a jQuery-like document
// from returned html, grabs all elements from the body, removes all "script-like" tags
// from the selection, and leverages the jQuery-style "text" function to extract text
func HtmlTextContent(res *http.Response) (string, error) {
	// response comes from remote server, must be closed when we're done
	defer res.Body.Close()
	// goquery is the package that does most of the heavy lifting here
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return "", err
	}

	// remove unwanted elements
	doc.Find("script,style,noscript,canvas,meta").Remove()

	return doc.Find("body").Text(), nil
}
