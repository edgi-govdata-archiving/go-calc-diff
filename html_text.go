package main

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

var (
	textElements       = strings.Join([]string{"blockquote", "dd", "dl", "dt", "figcaption", "figure", "hr", "li", "main", "ol", "p", "pre", "ul"}, ",")
	inlineTextElements = strings.Join([]string{"a", "abbr", "b", "bdi", "bdo", "br", "cite", "code", "data", "dfn", "em", "i", "kbd", "mark", "q", "rp", "rt", "rtc", "ruby", "s", "samp", "small", "span", "strong", "sub", "sup", "time", "u", "var", "wbr"}, ",")
	tableElements      = strings.Join([]string{"caption", "col", "colgroup", "table", "tbody", "td", "tfoot", "th", "thead", "tr"}, ",")
	formElements       = strings.Join([]string{"button", "datalist", "fieldset", "form", "input", "label", "legend", "meter", "optgroup", "option", "output", "progress", "select", "textarea"}, ",")
	selector           = strings.Join([]string{textElements, inlineTextElements, tableElements, formElements}, ",")
)

func HtmlTextContent(res *http.Response) (string, error) {
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return "", err
	}

	added := 0
	text := ""

	doc.Find(selector).Each(func(i int, el *goquery.Selection) {
		t := el.Text()
		if t != "" {
			if added == 0 {
				text = el.Text()
			} else {
				text += " " + el.Text()
			}
			added++
		}
	})

	return strings.TrimSpace(text), nil
}
