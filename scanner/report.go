package scanner

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
)

type Report struct {
	Vulnerabilities []Vulnerability
}

// GenerateReport creates a vulnerability report in JSON or HTML
func GenerateReport(vulnerabilities []Vulnerability, fileName, format string) error {
	report := Report{Vulnerabilities: vulnerabilities}

	if format == "json" {
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		return encoder.Encode(report)
	} else if format == "html" {
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		tmpl := `<html><body>
		<h1>Vulnerability Report</h1>
		<table border="1">
		<tr><th>Package</th><th>Version</th><th>Vulnerability</th><th>Severity</th><th>Description</th></tr>
		{{range .Vulnerabilities}}
		<tr>
		<td>{{.Package}}</td><td>{{.Version}}</td><td>{{.Vulnerability}}</td><td>{{.Severity}}</td><td>{{.Description}}</td>
		</tr>
		{{end}}
		</table></body></html>`

		t := template.Must(template.New("report").Parse(tmpl))
		return t.Execute(file, report)
	}
	return fmt.Errorf("unsupported format: %s", format)
}
