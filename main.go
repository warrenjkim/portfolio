package main

import (
	"portfolio/internal/generator"
)

func main() {
	generator := generator.NewGenerator("dist")
	generator.Build()
}
