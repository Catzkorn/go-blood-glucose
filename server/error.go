package server

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderError(err error, w http.ResponseWriter) {
	// Constructs information for template
	data := struct {
		Error error
	}{
		Error: err,
	}

	err = parsedErrorTemplate.Execute(w, data)
	if err != nil {
		fmt.Println("Failed to execute template:", err)
		return
	}

}

const errorTemplate = `<!DOCTYPE html>
<html>

<head>
	
	<meta charset="UTF-8">
	
</head>

<body>

<p><strong>There was an error: {{ .Error }}</strong></p>

<a href="/">Go Back</a>


</body>`

var parsedErrorTemplate = template.Must(template.New("error").Parse(errorTemplate))
