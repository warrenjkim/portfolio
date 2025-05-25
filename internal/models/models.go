package models

import (
	"time"
)

type Link struct {
	Url         string
	DisplayName string
	ImagePath   string
	AltName     string
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
	Files  []NoteFile
}

type NoteFile struct {
	Name      string
	PdfPath   string
	LatexPath string
}

type Courses struct {
	Ucla               []string
	IvcSaddlebackMajor []string
	IvcSaddlebackIgetc []string
}
