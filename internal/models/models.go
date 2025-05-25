package models

import (
	"html/template"
	"time"
)

type Link struct {
	Url         string
	DisplayName string
	ImagePath   string
	AltName     string
}

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

type ProjectPreview struct {
	Name       string
	Caption    string
	GitHubUrl  string
	CoverImage string
	StartDate  time.Time
	EndDate    *time.Time
	Skills     []Skill
}

type ProjectImage struct {
	Src     string
	Caption string
}

type Skill struct {
	Name      string
	ImagePath string
}

type Note struct {
	Course string
	Files      []NoteFile
}

type NoteFile struct {
	Name      string
	PdfPath   string
	LatexPath string
}

type Courses struct {
	Ucla []string
	IvcSaddlebackMajor []string
	IvcSaddlebackIgetc []string
}
