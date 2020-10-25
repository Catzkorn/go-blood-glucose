package server

import "html/template"

const indexTemplate = `<!DOCTYPE html>
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

{{ if .MonitorCreated }}


<div class="add_reading">
<form action="/add_reading" method="post">
<label for="upper">Add Reading:</label>
<input type="text" name="reading"><br>
<input type="submit" value="Submit">
</form>
</div>

{{end}}


{{range .Readings}}<div>{{ . }}</div>{{else}}<div><strong>no readings</strong></div>{{end}}



</body>`

var parsedIndexTemplate = template.Must(template.New("index").Parse(indexTemplate))
