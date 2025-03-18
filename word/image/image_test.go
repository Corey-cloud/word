package image

import (
	"github.com/carmel/gooxml/document"
	"testing"
)

func TestImage_InsertImage(t *testing.T) {
	doc := document.New()
	image, err := NewImage()
	if err != nil {
		t.Error("SaveToFile err: ", err)
	}
	path := "image.png"
	name := "趋势图"
	if err := image.InsertImage(doc, path, name); err != nil {
		t.Error("InsertImage err: ", err.Error())
	}
	err = doc.SaveToFile("image.docx")
	if err != nil {
		t.Error("SaveToFile err: ", err.Error())
	}
}
