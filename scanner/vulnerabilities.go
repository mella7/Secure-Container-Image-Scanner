package scanner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Vulnerability struct {
	Package        string `json:"package"`
	Version        string `json:"version"`
	Vulnerability  string `json:"vulnerability"`
	Severity       string `json:"severity"`
	Description    string `json:"description"`
}

// CheckVulnerabilities queries a vulnerability database for image layers
func CheckVulnerabilities(layers []Layer) ([]Vulnerability, error) {
	client := resty.New()
	vulnerabilities := []Vulnerability{}

	for _, layer := range layers {
		resp, err := client.R().
			SetQueryParam("digest", layer.Digest).
			Get("https://example-vulnerability-api.com/layers")
		if err != nil {
			return nil, err
		}

		if resp.StatusCode() != http.StatusOK {
			return nil, fmt.Errorf("failed to fetch data for layer %s", layer.Digest)
		}

		var layerVulnerabilities []Vulnerability
		err = json.Unmarshal(resp.Body(), &layerVulnerabilities)
		if err != nil {
			return nil, err
		}

		vulnerabilities = append(vulnerabilities, layerVulnerabilities...)
	}

	return vulnerabilities, nil
}
