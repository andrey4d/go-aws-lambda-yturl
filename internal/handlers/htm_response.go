/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"bytes"
	"html/template"

	"github.com/aws/aws-lambda-go/events"
)

func HtmlResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	response := events.APIGatewayProxyResponse{
		Headers: map[string]string{"Content-Type": "text/html"},
	}
	response.StatusCode = status

	tmpl, err := getHtmlTemplate()
	if err != nil {
		return nil, err
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, body); err != nil {
		return nil, err
	}

	response.Body = tpl.String()
	return &response, nil
}

func getHtmlTemplate() (*template.Template, error) {
	tmp := `<!DOCTYPE html>
	<html>
	<head>
		<meta charset='utf-8'>
		<meta http-equiv='X-UA-Compatible' content='IE=edge'>
		<title>{{ .Title }}</title>
		<meta name='viewport' content='width=device-width, initial-scale=1'>
		<style>
    	footer { text-align: center; padding: 3px; background-color:  #0047ab; color: white;}
        header { text-align: center; padding: 3px; background-color: #0047ab; color: white;}
		</style>

	</head>
	<header> <h1>{{ .Title}}</h1>  </header>
	<body>
		<h1>Author: {{ .Author }} </h1>
		<p>Description:</p>
		<details>
		<summary> {{ .Description }}</summary>
		</details>
		<ul>
		
		{{ range .Formats}}
		<li>
			<a href="{{.URL}}">Video={{.QualityLabel}} Audio={{.AudioQuality}} Audio channels={{.AudioChannels}} Type={{.MimeType}}</a>
		</li>
		{{ end }}

		</ul>
	</body>
	<footer>
    <p><a style="color: #ffd700 " href="mailto:andrey4d.dev@gmail.com">andrey4d.dev@gmail.com</a></p>
</footer>

	</html>`
	tmpl, err := template.New("base").Parse(tmp)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
