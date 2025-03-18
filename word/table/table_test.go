package table

import (
	"github.com/carmel/gooxml/document"
	"testing"
)

func TestTable_SetBackgroundColor(t *testing.T) {
	doc := document.New()
	cell := doc.AddTable().AddRow().AddCell()
	cell.AddParagraph().AddRun().AddText("背景-自定义-测试紫")
	NewTableSet().SetBackgroundColor(&cell, "6b4f9a")
	err := doc.SaveToFile("table_cell_purple.docx")
	if err != nil {
		t.Error("SaveToFile err: ", err.Error())
	}
}

func TestTable_SetBackgroundColorGray(t *testing.T) {
	doc := document.New()
	cell := doc.AddTable().AddRow().AddCell()
	cell.AddParagraph().AddRun().AddText("背景-灰")
	NewTableSet().SetBackgroundColorGray(&cell)
	err := doc.SaveToFile("table_cell_gray.docx")
	if err != nil {
		t.Error("SaveToFile err: ", err.Error())
	}
}

func TestTable_SetBackgroundColorBlue(t *testing.T) {
	doc := document.New()
	cell := doc.AddTable().AddRow().AddCell()
	cell.AddParagraph().AddRun().AddText("背景-蓝")
	NewTableSet().SetBackgroundColorBlue(&cell)
	err := doc.SaveToFile("table_cell_blue.docx")
	if err != nil {
		t.Error("SaveToFile err: ", err.Error())
	}
}

func TestTable_SetBackgroundColorGreen(t *testing.T) {
	doc := document.New()
	cell := doc.AddTable().AddRow().AddCell()
	cell.AddParagraph().AddRun().AddText("背景-绿")
	NewTableSet().SetBackgroundColorGreen(&cell)
	err := doc.SaveToFile("table_cell_green.docx")
	if err != nil {
		t.Error("SaveToFile err: ", err.Error())
	}
}

func TestTable_SetBackgroundColorYellow(t *testing.T) {
	doc := document.New()
	cell := doc.AddTable().AddRow().AddCell()
	cell.AddParagraph().AddRun().AddText("背景-黄")
	NewTableSet().SetBackgroundColorYellow(&cell)
	err := doc.SaveToFile("table_cell_yellow.docx")
	if err != nil {
		t.Error("SaveToFile err: ", err.Error())
	}
}

func TestTable_SetBackgroundColorOrange(t *testing.T) {
	doc := document.New()
	cell := doc.AddTable().AddRow().AddCell()
	cell.AddParagraph().AddRun().AddText("背景-橙")
	NewTableSet().SetBackgroundColorOrange(&cell)
	err := doc.SaveToFile("table_cell_orange.docx")
	if err != nil {
		t.Error("SaveToFile err: ", err.Error())
	}
}

func TestTable_SetBackgroundColorRed(t *testing.T) {
	doc := document.New()
	cell := doc.AddTable().AddRow().AddCell()
	cell.AddParagraph().AddRun().AddText("背景-红")
	NewTableSet().SetBackgroundColorRed(&cell)
	err := doc.SaveToFile("table_cell_red.docx")
	if err != nil {
		t.Error("SaveToFile err: ", err.Error())
	}
}
