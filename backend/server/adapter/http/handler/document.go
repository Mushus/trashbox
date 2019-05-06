package handler

import (
	"net/http"

	"golang.org/x/xerrors"

	"github.com/Mushus/trashbox/backend/server/app/property"

	"github.com/Mushus/trashbox/backend/server/adapter/http/template"
	"github.com/Mushus/trashbox/backend/server/app/document"
)

// GetIndex is a handler show index of webpage
func (h Handler) GetIndex(c Context) error {
	return c.String(http.StatusOK, "it's works!")
}

// GetDocument is a handler of get document
func (h Handler) GetDocument(c Context) error {
	title := c.Param("title")

	_, err := h.document.Get(title)
	if xerrors.Is(err, document.ErrDocumentNotFound) {
		if !c.IsLoggedIn {
			return c.String(http.StatusNotFound, "document not found")
		}
		return c.Render(http.StatusOK, template.TmplEdit, template.EditView{
			Title: title,
		})
	}
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "") // TODO:
}

// PutDocumentParam is a parameter used as PutDocument handler
type PutDocumentParam struct {
	Content string `json:"content" validate:"required"`
}

// PutDocument is a handler to save document
func (h Handler) PutDocument(c Context) error {
	title := c.Param("title")

	prm := PutDocumentParam{}
	if err := c.Bind(&prm); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	if err := c.Validate(prm); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	doc := property.Document{
		Title:   title,
		Content: prm.Content,
	}

	if err := h.document.Put(&doc); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, struct{}{})
}
