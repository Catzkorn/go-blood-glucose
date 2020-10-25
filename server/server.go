package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/Catzkorn/go-blood-glucose/monitor"
	"github.com/shopspring/decimal"
)

// Server defines a server
type Server struct {
	monitor monitor.Monitor
}

// Handle implements a standard library http.Handler
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/update_monitor":
		monitor, err := monitor.New(r.FormValue("upper"), r.FormValue("lower"))

		if err != nil {
			renderError(err, w)
			return
		}
		s.monitor = monitor

		http.Redirect(w, r, "/", http.StatusFound)
		return

	case "/add_reading":
		err := s.monitor.AddReading(r.FormValue("reading"))

		if err != nil {
			renderError(err, w)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	reqbytes, err := httputil.DumpRequest(r, true)

	if err == nil {
		fmt.Println(string(reqbytes))
	}

	// Constructs information for template
	data := struct {
		Readings []decimal.Decimal
	}{
		Readings: s.monitor.Readings(),
	}

	err = parsedIndexTemplate.Execute(w, data)
	if err != nil {
		fmt.Println("Failed to execute template:", err)
		return
	}
}
