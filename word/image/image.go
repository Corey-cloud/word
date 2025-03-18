package image

import (
	"fmt"
	"github.com/carmel/gooxml/common"
	"github.com/carmel/gooxml/document"
	"github.com/carmel/gooxml/measurement"
	"github.com/carmel/gooxml/schema/soo/wml"
)

type Image struct {
	width  int
	height int
}

func NewImage(opts ...int) (*Image, error) {
	if len(opts) > 2 {
		return nil, fmt.Errorf("参数错误")
	}
	img := &Image{}
	// 根据参数数量初始化 width 和 height
	switch len(opts) {
	case 0:
		// 默认值
		img.width = 416
		img.height = 234
	case 1:
		img.width = opts[0]
		img.height = opts[0]
	case 2:
		img.width = opts[0]
		img.height = opts[1]
	}
	return img, nil
}

func (i *Image) InsertImage(doc *document.Document, path, name string) error {
	p := doc.AddParagraph()
	p.Properties().SetAlignment(wml.ST_JcCenter)
	r := p.AddRun()
	r.Properties().SetSize(12 * measurement.Point)
	img, err := common.ImageFromFile(path)
	if err != nil {
		return err
	}
	imageRef, err := doc.AddImage(img)
	if err != nil {
		return err
	}
	anchored, err := r.AddDrawingInline(imageRef)
	if err != nil {
		return err
	}
	anchored.SetSize(measurement.Distance(i.width*measurement.Point), measurement.Distance(i.height*measurement.Point))
	r.AddBreak()
	r.AddText(name)
	return nil
}
