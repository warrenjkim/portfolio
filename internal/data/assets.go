package data

import (
	"portfolio/internal/models"
	"time"
)

var (
	skills = map[string]models.Skill{
		"C++":    {Name: "C++", ImagePath: "static/assets/skills/C++.png"},
		"C":      {Name: "C", ImagePath: "static/assets/skills/C.png"},
		"Go":     {Name: "Go", ImagePath: "static/assets/skills/Golang.png"},
		"Python": {Name: "Python", ImagePath: "static/assets/skills/Python.png"},
		"Git":    {Name: "Git", ImagePath: "static/assets/skills/Git.png"},
	}

	links = []models.Link{
		{
			Url:         "https://github.com/warrenjkim",
			DisplayName: "warrenjkim",
			ImagePath:   "static/assets/links/github.webp",
			AltName:     "GitHub",
		},
		{
			Url:         "https://linkedin.com/in/warren-kim",
			DisplayName: "warren-kim",
			ImagePath:   "static/assets/links/linkedin.webp",
			AltName:     "LinkedIn",
		},
		// {
		// 	Url:         "https://youtube.com/@warren_kim",
		// 	DisplayName: "@warren_kim",
		// 	ImagePath:   "static/assets/links/youtube.webp",
		// 	AltName:     "YouTube",
		// },
	}

	socials = links
	// append(links,
	// 		models.Link{
	// 			Url:         "https://open.spotify.com/artist/3uoZD1Z2Zcv2adealPuYpm?si=F4K44vtISWe9rUj6AKMEbg",
	// 			DisplayName: "",
	// 			ImagePath:   "static/assets/links/spotify.webp",
	// 			AltName:     "Spotify",
	// 		})
)

func LoadProjectPreviews() []models.ProjectPreview {
	parseDate := func(dateStr string) time.Time {
		t, _ := time.Parse("2006-01-02", dateStr)

		return t
	}

	parseDatePtr := func(dateStr string) *time.Time {
		if dateStr == "" {
			return nil
		}

		t := parseDate(dateStr)

		return &t
	}

	return []models.ProjectPreview{
		{
			Name:       "JSON Parser",
			Caption:    "A standards compliant JSON Parser in C++!",
			GitHubUrl:  "https://github.com/warrenjkim/json",
			CoverImage: "",
			StartDate:  parseDate("2024-08-01"),
			EndDate:    parseDatePtr(""),
			Skills: []models.Skill{
				skills["C++"],
			},
		},
		{
			Name:       "Memory Driver",
			Caption:    "A cache visualization system to analyze data patterns through a three-level CPU cache.",
			GitHubUrl:  "https://github.com/warrenjkim/memory-driver",
			CoverImage: "",
			StartDate:  parseDate("2023-10-01"),
			EndDate:    parseDatePtr("2023-11-01"),
			Skills: []models.Skill{
				skills["C++"],
			},
		},
	}
}

func LoadNotes() []models.Note {
	return []models.Note{
		{
			Course: "Ring Theory (MATH 110A)",
			Files: []models.NoteFile{
				{
					Name:      "Notes",
					PdfPath:   "static/assets/notes/MATH110A-notes.pdf",
					LatexPath: "static/assets/notes/MATH110A-notes.zip",
				},
			},
		},
		{
			Course: "Real Analysis (MATH 131A)",
			Files: []models.NoteFile{
				{
					Name:      "Midterm 1 Cheat Sheet",
					PdfPath:   "static/assets/notes/MATH131A-midterm-1-cheat-sheet.pdf",
					LatexPath: "static/assets/notes/MATH131A-midterm-1-cheat-sheet.zip",
				},
				{
					Name:      "Midterm 2 Cheat Sheet",
					PdfPath:   "static/assets/notes/MATH131A-midterm-2-cheat-sheet.pdf",
					LatexPath: "static/assets/notes/MATH131A-midterm-2-cheat-sheet.zip",
				},
				{
					Name:      "Final Cheat Sheet",
					PdfPath:   "static/assets/notes/MATH131A-final-cheat-sheet.pdf",
					LatexPath: "static/assets/notes/MATH131A-final-cheat-sheet.zip",
				},
			},
		},
		{
			Course: "Operating Systems (COM SCI 111)",
			Files: []models.NoteFile{
				{
					Name:      "Notes",
					PdfPath:   "static/assets/notes/CS111-notes.pdf",
					LatexPath: "static/assets/notes/CS111-notes.zip",
				},
			},
		},
		{
			Course: "Software Construction (COM SCI 35L)",
			Files: []models.NoteFile{
				{
					Name:      "Notes",
					PdfPath:   "static/assets/notes/CS35L-notes.pdf",
					LatexPath: "static/assets/notes/CS35L-notes.zip",
				},
			},
		},
		{
			Course: "Programming Languages (COM SCI 131)*",
			Files: []models.NoteFile{
				{
					Name:      "Midterm Notes",
					PdfPath:   "static/assets/notes/CS131-midterm-notes.pdf",
					LatexPath: "static/assets/notes/CS131-midterm-notes.zip",
				},
				{
					Name:      "Final Notes",
					PdfPath:   "static/assets/notes/CS131-final-notes.pdf",
					LatexPath: "static/assets/notes/CS131-final-notes.zip",
				},
			},
		},
		{
			Course: "Data Management Systems (COM SCI 143)",
			Files: []models.NoteFile{
				{
					Name:      "Cheat Sheet",
					PdfPath:   "static/assets/notes/CS143-cheat-sheet.pdf",
					LatexPath: "static/assets/notes/CS143-cheat-sheet.zip",
				},
				{
					Name:      "Notes",
					PdfPath:   "static/assets/notes/CS143-notes.pdf",
					LatexPath: "static/assets/notes/CS143-notes.zip",
				},
			},
		},
		{
			Course: "Deep Learning and Neural Networks (EC ENGR C147)",
			Files: []models.NoteFile{
				{
					Name:      "Cheat Sheet",
					PdfPath:   "static/assets/notes/ECEC147-cheat-sheet.pdf",
					LatexPath: "static/assets/notes/ECEC147-cheat-sheet.zip",
				},
			},
		},
		{
			Course: "Introduction to Finance and Marketing for Engineers (ENGR 111)",
			Files: []models.NoteFile{
				{
					Name:      "Midterm Cheat Sheet",
					PdfPath:   "static/assets/notes/ENGR111-midterm-cheat-sheet.pdf",
					LatexPath: "static/assets/notes/ENGR111-midterm-cheat-sheet.zip",
				},
				{
					Name:      "Final Cheat Sheet",
					PdfPath:   "static/assets/notes/ENGR111-final-cheat-sheet.pdf",
					LatexPath: "static/assets/notes/ENGR111-final-cheat-sheet.zip",
				},
			},
		},
	}
}

func LoadCourses() models.Courses {
	return models.Courses{
		Ucla: []string{
			"(Real) Analysis (MATH 131)",
			"Operating Systems (COM SCI 111)",
			"Software Construction (COM SCI 35L)",
			"Software Engineering (COM SCI 130)",
			"Programming Languages (COM SCI 131)",
			"Compiler Construction (COM SCI 132)",
			"Data Management Systems (COM SCI 143)",
			"(Proof-based) Linear Algebra (MATH 115A)",
			"Computer Systems Architecture (ECE M116C)",
			"Computer Network Fundamentals (COM SCI 118)",
			"Logic Design of Digital Systems (COM SCI M51A)",
			"Introduction to Computer Graphics (COM SCI 174A)",
			"Deep Learning and Neural Networks (EC ENGR C147)",
			"Fundamentals of Artificial Intelligence (COM SCI 161)",
			"Introductory Digital Design Laboratory (COM SCI M152A)",
			"Introduction to Algorithms and Complexity (COM SCI 180)",
			"Experimental Courses in Engineering Ethics (ENGR 188EW)",
			"Introduction to Finance and Marketing for Engineers (ENGR 111)",
			"Introduction to Probability and Statistics for Engineers (C&EE 110)",
			"Introduction to Formal Languages and Automata Theory (COM SCI 181)",
		},
		IvcSaddlebackMajor: []string{
			"C Programming (CS 36)",
			"Java Programming (CS 38)",
			"Data Structures and Algorithms (CS 1D)",
			"General Physics I/II/III (PHYS 4A/B/C)",
			"Introduction to Computer Systems (CS 1)",
			"Introduction to Linear Algebra (MATH 26)",
			"Elementary Differential Equations (MATH 24)",
			"(Calculus III) Analytical Geometry (MATH 4A)",
			"Discrete Mathematics I/II (CS 6A/B or MATH 30/31)",
			"Introduction to Computer Science I/II/III (CS 1A/B/C)",
			"Computer Organization and Assembly Language I/II (CS 40A/B)",
		},
		IvcSaddlebackIgetc: []string{
			"College Writing II (WR 2)",
			"Music Appreciation (MUS 20)",
			"Life Sciences/Lab (BIO 1/1L)",
			"Honors Film & US Culture (FILM 72H)",
			"Communication Fundamentals (COMM 1H)",
			"Honors Introduction to Psychology (PSYCH 1H)",
			"Honors Academic, Career, Life Counseling (COUN 6H)",
		},
	}
}
