package config

import (
	"html/template"
	"portfolio/internal/models"
)

var (
	Header = models.Header{
		Navigation: []models.Link{
			{
				Url:         "/about",
				DisplayName: "About",
			},
			{
				Url:         "/projects",
				DisplayName: "Projects",
			},
			{
				Url:         "/notes",
				DisplayName: "Notes",
			},
		},
		Socials: data.LoadLinks(true),
	}

	Footer = models.Footer{
		Copyright: template.HTML("&copy; 2025 Warren Kim"),
		Socials:   data.LoadLinks(true),
	}
)
