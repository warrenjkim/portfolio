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
	Header  Header
	Footer  Footer
	Name    string
	Links   []Link
	About   string
	Facts   []template.HTML
	Hobbies []string
}

// TODO(Notes, Courses, Projects)
