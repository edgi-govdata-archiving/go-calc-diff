/*
	Coverage is a service for mapping an archiving surface area, and tracking
	the amount of that surface area that any number of archives have covered
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	// cfg is the global configuration for the server. It's read in at startup from
	// the config.json file and enviornment variables, see config.go for more info.
	cfg *config

	// When was the last alert sent out?
	// Use this value to avoid bombing alerts
	lastAlertSent *time.Time

	// log output
	logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
)

func main() {
	var err error
	// setup config
	cfg, err = initConfig(os.Getenv("GOLANG_ENV"))
	if err != nil {
		// panic if the server is missing a vital configuration detail
		logger.Fatal(fmt.Errorf("server configuration error: %s", err.Error()))
	}

	s := &http.Server{}
	m := http.NewServeMux()
	m.Handle("/", middleware(DiffHandler))
	m.HandleFunc("/.well-known/acme-challenge/", CertbotHandler)

	// connect mux to server
	s.Handler = m

	// fire it up!
	fmt.Println("starting server on port", cfg.Port)
	// start server wrapped in a log.Fatal b/c http.ListenAndServe will not
	// return unless there's an error, which would be a program crash
	logger.Fatal(StartServer(cfg, s))
}
