package renderer

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

type (
	// TemplateMap 複数のテンプレートをまとめたもの
	TemplateMap map[string]*template.Template
	// Renderer テンプレートエンジン
	Renderer struct {
		templates TemplateMap
	}
)

// ExecuteTemplate テンプレートを実行する
func (t TemplateMap) ExecuteTemplate(w io.Writer, name string, data interface{}) error {
	if tmpl, ok := t[name]; ok {
		return tmpl.ExecuteTemplate(w, "layout", data)
	}
	return xerrors.Errorf("template %v is not found", name)
}

// ProvideRenderer 新しいテンプレートを作成する
func ProvideRenderer(tmpl TemplateMap) *Renderer {
	return &Renderer{
		templates: tmpl,
	}
}

// Render レンダリングを行います
func (t *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
