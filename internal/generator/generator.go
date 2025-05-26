package generator

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"portfolio/internal/data"
)

// TODO(Notes, Courses, Projects)

type Generator struct {
	templates *template.Template
	outputDir string
}

func NewGenerator(outputDir string) *Generator {
	templates := template.Must(template.ParseGlob("templates/**/*.tmpl"))
	for _, t := range templates.Templates() {
		log.Printf("Loaded template: %s", t.Name())
	}

	return &Generator{
		templates: templates,
		outputDir: outputDir,
	}
}

func (g *Generator) Build() {
	err := os.MkdirAll(g.outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to create output directory %s: %v", g.outputDir, err)
	}

	pages := data.LoadPages()

	for _, page := range pages {
		log.Printf("generating %s...\n", page.Path)
		g.generatePage(page.Path, page.Template, page.Context)
		log.Printf("done generating %s!\n", page.Path)
	}
}

func (g *Generator) generatePage(path string, name string, context interface{}) {
	outPath := filepath.Join(g.outputDir, path)
	file, err := os.Create(outPath)
	if err != nil {
		log.Fatalf("failed to create file %s: %v", outPath, err)
	}
	defer file.Close()

	if err := g.templates.ExecuteTemplate(file, name, context); err != nil {
		log.Fatalf("failed to execute template %s: %v", outPath, err)
	}
}
