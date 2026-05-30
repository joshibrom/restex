// Package render handles rendering LaTeX files.
package render

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/joshibrom/restex/internal/latex"
	"github.com/joshibrom/restex/internal/model"
)

func Render(resume *model.Resume) (string, error) {
	funcMap := template.FuncMap{
		"join":      strings.Join,
		"prepare":   latex.Prepare,
		"phoneHref": resume.Contact.Phone.HrefValue,
		"phoneText": resume.Contact.Phone.TextValue,
	}
	templ := template.Must(template.New("latex").Funcs(funcMap).ParseFiles(
		"templates/resume.tex.tmpl",
		"templates/sections/preamble.tex.tmpl",
		"templates/sections/heading.tex.tmpl",
		"templates/sections/education.tex.tmpl",
		"templates/sections/experience.tex.tmpl",
		"templates/sections/project.tex.tmpl",
		"templates/sections/skill.tex.tmpl",
		"templates/sections/service.tex.tmpl",
	))

	var output bytes.Buffer

	err := templ.ExecuteTemplate(&output, "latex", resume)
	if err != nil {
		return "", nil
	}

	return output.String(), nil
}
