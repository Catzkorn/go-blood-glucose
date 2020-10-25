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
	monitor *monitor.Monitor
}

// Handle implements a standard library http.Handler
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	reqbytes, err := httputil.DumpRequest(r, true)

	if err == nil {
		fmt.Println(string(reqbytes))
	}

	switch r.URL.Path {
	case "/update_monitor":
		monitor, err := monitor.New(r.FormValue("upper"), r.FormValue("lower"))

		if err != nil {
			renderError(err, w)
			return
		}
		s.monitor = &monitor

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

	// This is an easy way to reverse a slice in go https://github.com/golang/go/wiki/SliceTricks#reversing
	readings := s.monitor.Readings()
	for i := len(readings)/2 - 1; i >= 0; i-- {
		opp := len(readings) - 1 - i
		readings[i], readings[opp] = readings[opp], readings[i]
	}

	// Constructs information for template
	data := struct {
		Readings       []decimal.Decimal
		MonitorCreated bool
	}{

		Readings:       readings,
		MonitorCreated: s.monitor != nil,
	}

	err = parsedIndexTemplate.Execute(w, data)
	if err != nil {
		fmt.Println("Failed to execute template:", err)
		return
	}
}
