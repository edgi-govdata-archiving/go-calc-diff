package main

import (
	"fmt"
	"github.com/qri-io/go-diff/diffmatchpatch"
	"io"
	"net/http"
	"strconv"
	"text/template"
)

var (
	// TODO - currently templates are being rendered with the "text/template" package
	// as a quick fix for rendering "dangerous" html. This should be replaced
	// with the html/template package
	templates = template.Must(template.ParseFiles("views/diff.html"))
)

// DiffHandler routes requests to the right func
func DiffHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		EmptyOkHandler(w, r)
	case "GET", "POST":
		CalcDiffHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

// CalcDiffHandler takes lots of query params to calc diffs
func CalcDiffHandler(w http.ResponseWriter, r *http.Request) {
	var diffs []diffmatchpatch.Diff
	var err error

	// url values as a & b query params or POST form values
	a := r.FormValue("a")
	b := r.FormValue("b")
	// snag output format from request
	format := r.FormValue("format")

	// do we want to diff html text?
	htmlText := r.FormValue("html_text")

	// create & configure a differ from the request
	dmp := NewDifferFromRequest(r)

	if a == "" || b == "" {
		writeErrResponse(w, http.StatusBadRequest, fmt.Errorf("please supply both a & b urls"))
		return
	}

	// compare a & b to generate diffs
	if htmlText == "true" {
		diffs, err = HtmlTextDiff(dmp, a, b)
	} else {
		diffs, err = HtmlDiff(dmp, a, b)
	}

	if err != nil {
		writeErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	// generate response
	switch format {
	case "json":
		writeResponse(w, diffs)
	case "htmlraw":
		io.WriteString(w, dmp.DiffPrettyHtml(diffs))
	case "text":
		w.Header().Add("Content-Type", "text")
		io.WriteString(w, dmp.DiffPrettyText(diffs))
	default:
		// default to "html"
		renderTemplate(w, "diff.html", dmp.DiffPrettyHtml(diffs))
	}
}

// HealthCheckHandler is a basic "hey I'm fine" for load balancers
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{ "status" : 200 }`))
}

// EmptyOkHandler is an empty 200 response, often used
// for OPTIONS requests that responds with headers set in addCorsHeaders
func EmptyOkHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// CertbotHandler pipes the certbot response for manual certificate generation
func CertbotHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, cfg.CertbotResponse)
}

// NotFoundHandler is a basic JSON 404
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{ "status" :  "not found" }`))
}

// renderTemplate renders a template with the values of cfg.TemplateData
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// reqParamInt extracts an int from an http request based on a key.
// it'll check query params and POST form values
func reqParamInt(key string, r *http.Request) (int, error) {
	i, err := strconv.ParseInt(r.FormValue(key), 10, 0)
	return int(i), err
}

// reqParamInt extracts a float64 from an http request based on a key.
// it'll check query params and POST form values
func reqParamFloat(key string, r *http.Request) (float64, error) {
	return strconv.ParseFloat(r.FormValue(key), 10)
}

// reqParamInt extracts a bool from an http request based on a key.
// it'll check query params and POST form values
func reqParamBool(key string, r *http.Request) (bool, error) {
	return strconv.ParseBool(r.FormValue(key))
}
