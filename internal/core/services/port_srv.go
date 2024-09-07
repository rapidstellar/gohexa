package services

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/rapidstellar/gohexa/internal/core/domain"
)

// GeneratePortsFile implements ports.IGeneratorService.
func (g *GeneratorServiceImpls) GeneratePortsFile(dir string) {
	// Default to current directory if not provided
	if dir == "" {
		dir = "./internal/core/ports"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("Error creating directories: %v\n", err)
			return
		}
	}
	// Prepare the data for template rendering
	data := domain.PortFlagDomain{
		FeatureName: g.flag.FeatureName,
		ProjectName: g.flag.ProjectName,
		IDType:      "uint", // Default to uint, can be changed to string if UUID is used
	}

	// Parse and execute the template
	tmpl, err := template.New("ports").Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
	}).Parse(domain.PortsTemplate)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	// Create the output file path
	fileName := fmt.Sprintf("%s_ports.go", strings.ToLower(g.flag.FeatureName))
	filePath := filepath.Join(dir, fileName)

	// Create the output file
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	// Execute the template and write to the file
	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
	} else {
		fmt.Printf("Ports file '%s' created successfully!\n", filePath)
	}
}
