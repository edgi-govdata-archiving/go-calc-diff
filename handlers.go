package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"text/template"
)

var (
	templates = template.Must(template.ParseFiles("views/diff.html"))
)

// renderTemplate renders a template with the values of cfg.TemplateData
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func reqParamInt(key string, r *http.Request) (int, error) {
	i, err := strconv.ParseInt(r.FormValue(key), 10, 0)
	return int(i), err
}

func reqParamBool(key string, r *http.Request) (bool, error) {
	return strconv.ParseBool(r.FormValue(key))
}

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

func CalcDiffHandler(w http.ResponseWriter, r *http.Request) {
	a := r.FormValue("a")
	b := r.FormValue("b")
	format := r.FormValue("format")
	if format == "" {
		format = "html"
	}

	if a == "" || b == "" {
		writeErrResponse(w, http.StatusBadRequest, fmt.Errorf("please supply both a & b urls"))
		return
	}

	aBody, err := FetchUrl(a)
	if err != nil {
		writeErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	bBody, err := FetchUrl(b)
	if err != nil {
		writeErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	diffs := CalcDiff(string(aBody), string(bBody))
	// io.WriteString(w, diff)
	switch format {
	case "json":
		writeResponse(w, diffs)
	case "html":
		renderTemplate(w, "diff.html", dmp.DiffPrettyHtml(diffs))
	case "htmlraw":
		io.WriteString(w, dmp.DiffPrettyHtml(diffs))
	case "text":
		io.WriteString(w, dmp.DiffPrettyText(diffs))
	}
}

// HealthCheckHandler is a basic "hey I'm fine" for load balancers & co
// TODO - add Database connection & proper configuration checks here for more accurate
// health reporting
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

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{ "status" :  "not found" }`))
}
