package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/Catzkorn/go-blood-glucose/monitor"
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
			fmt.Println("failed to initialise monitor:", err)
			return
		}
		s.monitor = monitor

		http.Redirect(w, r, "/", http.StatusFound)
		return

	case "/add_reading":
		err := s.monitor.AddReading(r.FormValue("reading"))

		if err != nil {
			fmt.Println("Failed to read user input:", err)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	reqbytes, err := httputil.DumpRequest(r, true)

	if err == nil {
		fmt.Println(string(reqbytes))
	}

	w.Write([]byte(
		`<!DOCTYPE html>
	<html>
	
	<head>
		
		<meta charset="UTF-8">
		
	</head>
	
	<body>


	<p>Blood Glucose Monitoring Site</p>

	<div class="add_limits">
	<form action="/update_monitor" method="post">
	<label for="upper">Upper Limit:</label>
	<input type="text" name="upper"><br>
	<label for="lower">Lower Limit:</label>
  <input type="text" name="lower"><br>
  <input type="submit" value="Submit">
  </form>
</div>

<div class="add_reading">
	<form action="/add_reading" method="post">
	<label for="upper">Add Reading:</label>
	<input type="text" name="reading"><br>
  <input type="submit" value="Submit">
  </form>
</div>


</body>`),
	)
}
