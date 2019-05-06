package document

import (
	"github.com/Mushus/trashbox/backend/server/app/property"
)

// Document is a document that is not parsed
type Document struct {
	title   string
	content string
}

func NewDocument(prop *property.Document) *Document {
	if prop == nil {
		return nil
	}
	return &Document{
		title:   prop.Title,
		content: prop.Content,
	}
}

func (d *Document) ToProp() *property.Document {
	if d == nil {
		return nil
	}
	return &property.Document{
		Title:   d.title,
		Content: d.content,
	}
}

func (d Document) Title() string {
	return d.title
}

func (d Document) Content() string {
	return d.content
}
