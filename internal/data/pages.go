package data

import (
	"html/template"
	"portfolio/internal/models"
)

var (
	header = models.Header{
		Navigation: []models.Link{
			{
				Url:         "/projects",
				DisplayName: "Projects",
			},
			{
				Url:         "/notes",
				DisplayName: "Notes",
			},
		},
		Socials: socials,
	}

	footer = models.Footer{
		Copyright: template.HTML("&copy; 2025 Warren Kim"),
		Socials:   socials,
	}
)

func loadIndex() *models.Index {
	return &models.Index{
		Header:   header,
		Footer:   footer,
		Headshot: "static/assets/headshot.jpg",
		Name:     "Warren Kim",
		Links:    links,
		About:    "",
		Facts: []template.HTML{
			"I graduated from UCLA.",
			"My favorite programming language is C++.",
			"I typically typeset everything in " +
				"<span class=\"latex\">L<sup>a</sup>T<sub>e</sub>X</span> since I think " +
				"it looks neat.",
			"I gave a <a href=\"https://github.com/warrenjkim/rbtree-lecture\"><i>guest " +
				"lecture</i></a> on Red-Black trees for Professor Carey Nachenberg's data " +
				"structures class!",
			"My text editor of choice is Neovim. It used to be Emacs, which I still " +
				"use occasionally, but it couldn&apos;t handle large files.",
		},
		Hobbies: []template.HTML{
			"I enjoy building keyboards. My current keyboard is a 60% with linear " +
			"switches.",
			"I love to climb! I have been climbing for a little over 2 years now, " +
				"climbing up to and including a V10 indoors and V8 outdoors. I have only " +
				"climbed outdoors twice (both times to Joshua Tree), so I don&apos;t " +
				"have too many outdoor projects. Some notable climbs I&apos;ve sent are " +
				"Chili Sauce (V7), Strawberry Contraceptives (V7) and Mulligan Variation (V8). " +
				"Check out my climbing journey " +
				"<a href=\"https://instagram.com/half_crimped\"><i>here!</i></a>",
			"I enjoy playing the guitar and making music. I have written a couple of " +
				"<a href=\"https://open.spotify.com/artist/3uoZD1Z2Zcv2adealPuYpm?si=F4K44vtISWe9rUj6AKMEbg\"><i>songs</i></a> " +
				"and arranged several covers, some of which I have " +
				"<a href=\"https://youtube.com/@warren_kim\"><i>filmed.</i></a>",
		},
	}
}

func LoadPages() []models.Page {
	return []models.Page{
		{
			Path:     "index.html",
			Template: "index",
			Context:  loadIndex(),
		},
	}
}
