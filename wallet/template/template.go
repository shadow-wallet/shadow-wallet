package template

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Template struct {
	content string
}

func New(htmlPath string, cssPath ...string) (*Template, error) {
	var style []string
	html, err := os.ReadFile(htmlPath)
	if err != nil {
		return nil, err
	}
	for _, f := range cssPath {
		css, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}
		style = append(style, string(css))
	}
	return &Template{
		content: fmt.Sprintf("%s\n<style>%s</style>", html, strings.Join(style, "\n")),
	}, nil
}

func (t *Template) Execute(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = fmt.Fprint(w, t.content)
}
