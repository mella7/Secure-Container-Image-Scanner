package main

import (
	"fmt"
	"os"
	"secure-container-image-scanner/scanner"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: secure-container-image-scanner <image-name>")
		os.Exit(1)
	}

	imageName := os.Args[1]
	color.Green("Scanning image: %s\n", imageName)

	// Pull the image
	err := scanner.PullDockerImage(imageName)
	if err != nil {
		color.Red("Failed to pull image: %v\n", err)
		os.Exit(1)
	}

	// Parse image layers
	layers, err := scanner.ParseImageLayers(imageName)
	if err != nil {
		color.Red("Failed to parse image layers: %v\n", err)
		os.Exit(1)
	}

	// Fetch vulnerabilities
	vulnerabilities, err := scanner.CheckVulnerabilities(layers)
	if err != nil {
		color.Red("Failed to check vulnerabilities: %v\n", err)
		os.Exit(1)
	}

	// Generate report
	err = scanner.GenerateReport(vulnerabilities, "report.json", "json")
	if err != nil {
		color.Red("Failed to generate report: %v\n", err)
		os.Exit(1)
	}

	color.Green("Scan complete! Report saved to report.json\n")
}
