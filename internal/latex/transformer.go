// Package latex handles transformations/sanitizations of content prior to use
// in LaTeX source documents.
package latex

import (
	"bytes"
	"strings"
)

func Prepare(s string) string {
	return replaceBold(escapeText(s))
}

func escapeText(s string) string {
	return strings.NewReplacer(
		`\\`, `\textbackslash{}`,
		`{`, `\{`,
		`}`, `\}`,
		`#`, `\#`,
		`$`, `\$`,
		`%`, `\%`,
		`&`, `\&`,
		`_`, `\_`,
		`^`, `\textasciicircum{}`,
		`~`, `\textasciitilde{}`,
	).Replace(s)
}

func replaceBold(s string) string {
	insideBold := false

	var buf bytes.Buffer

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == '*' && s[i+1] == '*' {
			if !insideBold {
				buf.WriteString(`\textbf{`)
			} else {
				buf.WriteString(`}`)
			}
			insideBold = !insideBold
			i++
		} else {
			buf.WriteByte(s[i])
		}
	}

	return buf.String()
}
