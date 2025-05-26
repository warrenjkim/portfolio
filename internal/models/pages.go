package models

import (
	"html/template"
)

type Header struct {
	Navigation []Link
	Socials    []Link
}

type Footer struct {
	Copyright template.HTML
	Socials   []Link
}

type Index struct {
	Header   Header
	Footer   Footer
	Title    string
	Headshot string
	Name     string
	Links    []Link
	About    string
	Facts    []template.HTML
	Hobbies  []template.HTML
}

type Page struct {
	Path     string
	Template string
	Context  interface{}
}

// TODO(Notes, Courses, Projects)
