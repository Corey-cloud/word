package table

import (
	"github.com/carmel/gooxml/color"
	"github.com/carmel/gooxml/document"
	"github.com/carmel/gooxml/measurement"
	"github.com/carmel/gooxml/schema/soo/wml"
)

type Table struct {
	/*
		#e0e0e0 gray
		#00ccff blue
		#00ff00 green
		#ffff00 yellow
		#ff6600 orange
		#ff0000 red
	*/
	bgColor string
}

func NewTableSet() *Table {
	return &Table{}
}

func (t *Table) SetBackgroundColor(cell *document.Cell, bgRgb string) {
	shd := wml.NewCT_Shd()
	shd.ValAttr = wml.ST_ShdClear
	rgb := bgRgb
	shd.FillAttr = &wml.ST_HexColor{ST_HexColorRGB: &rgb}
	if cell.X().TcPr == nil {
		cell.X().TcPr = &wml.CT_TcPr{}
	}
	cell.X().TcPr.Shd = shd
}

func (t *Table) SetBackgroundColorGray(cell *document.Cell) {
	shd := wml.NewCT_Shd()
	shd.ValAttr = wml.ST_ShdClear
	rgb := "e0e0e0"
	shd.FillAttr = &wml.ST_HexColor{ST_HexColorRGB: &rgb}
	if cell.X().TcPr == nil {
		cell.X().TcPr = &wml.CT_TcPr{}
	}
	cell.X().TcPr.Shd = shd
}

func (t *Table) SetBackgroundColorBlue(cell *document.Cell) {
	shd := wml.NewCT_Shd()
	shd.ValAttr = wml.ST_ShdClear
	rgb := "00ccff"
	shd.FillAttr = &wml.ST_HexColor{ST_HexColorRGB: &rgb}
	if cell.X().TcPr == nil {
		cell.X().TcPr = &wml.CT_TcPr{}
	}
	cell.X().TcPr.Shd = shd
}

func (t *Table) SetBackgroundColorGreen(cell *document.Cell) {
	shd := wml.NewCT_Shd()
	shd.ValAttr = wml.ST_ShdClear
	rgb := "00ff00"
	shd.FillAttr = &wml.ST_HexColor{ST_HexColorRGB: &rgb}
	if cell.X().TcPr == nil {
		cell.X().TcPr = &wml.CT_TcPr{}
	}
	cell.X().TcPr.Shd = shd
}

func (t *Table) SetBackgroundColorYellow(cell *document.Cell) {
	shd := wml.NewCT_Shd()
	shd.ValAttr = wml.ST_ShdClear
	rgb := "ffff00"
	shd.FillAttr = &wml.ST_HexColor{ST_HexColorRGB: &rgb}
	if cell.X().TcPr == nil {
		cell.X().TcPr = &wml.CT_TcPr{}
	}
	cell.X().TcPr.Shd = shd
}

func (t *Table) SetBackgroundColorOrange(cell *document.Cell) {
	shd := wml.NewCT_Shd()
	shd.ValAttr = wml.ST_ShdClear
	rgb := "ff6600"
	shd.FillAttr = &wml.ST_HexColor{ST_HexColorRGB: &rgb}
	if cell.X().TcPr == nil {
		cell.X().TcPr = &wml.CT_TcPr{}
	}
	cell.X().TcPr.Shd = shd
}

func (t *Table) SetBackgroundColorRed(cell *document.Cell) {
	shd := wml.NewCT_Shd()
	shd.ValAttr = wml.ST_ShdClear
	rgb := "ff0000"
	shd.FillAttr = &wml.ST_HexColor{ST_HexColorRGB: &rgb}
	if cell.X().TcPr == nil {
		cell.X().TcPr = &wml.CT_TcPr{}
	}
	cell.X().TcPr.Shd = shd
}

func (t *Table) SetBackgroundColorSeparation(cell *document.Cell) {
	shd := wml.NewCT_Shd()
	shd.ValAttr = wml.ST_ShdClear
	rgb := "ddd9c4"
	shd.FillAttr = &wml.ST_HexColor{ST_HexColorRGB: &rgb}
	if cell.X().TcPr == nil {
		cell.X().TcPr = &wml.CT_TcPr{}
	}
	cell.X().TcPr.Shd = shd
}

func AddTable(doc *document.Document, stBorder int64, widthPercent float64) *document.Table {
	table := doc.AddTable()
	borders := table.Properties().Borders()
	if stBorder != 0 {
		borders.SetAll(wml.ST_BorderSingle, color.Auto, measurement.Point*1)
	}
	table.Properties().SetAlignment(wml.ST_JcTableCenter)
	if widthPercent != 0 {
		table.Properties().SetWidthPercent(widthPercent)
	}
	return &table
}

func AddRow(table *document.Table) *document.Row {
	row := table.AddRow()
	row.Properties().SetHeight(21, wml.ST_HeightRuleAtLeast)
	return &row
}

func AddCell(row *document.Row, width float64, verticalMerge, columnSpan int) *document.Cell {
	// verticalMerge 2 开始；1 合并
	cell := row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	if columnSpan != 0 {
		cell.Properties().SetColumnSpan(columnSpan)
	}
	if width != 0 {
		cell.Properties().SetWidth(measurement.Distance(width))
	}
	if verticalMerge == 2 {
		cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)
	} else if verticalMerge == 1 {
		cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
	}
	return &cell
}

func AddCellWidthPercent(row *document.Row, widthPercent float64, verticalMerge, columnSpan int) *document.Cell {
	// verticalMerge 2 开始；1 合并
	cell := row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	if columnSpan != 0 {
		cell.Properties().SetColumnSpan(columnSpan)
	}
	if widthPercent != 0 {
		cell.Properties().SetWidthPercent(widthPercent)
	}
	if verticalMerge == 2 {
		cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)
	} else if verticalMerge == 1 {
		cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
	}
	return &cell
}

func AddCellColor(cell *document.Cell, level string) {
	switch level {
	case "未运行":
		NewTableSet().SetBackgroundColorGray(cell)
	case "优秀":
		NewTableSet().SetBackgroundColorBlue(cell)
	case "良好":
		NewTableSet().SetBackgroundColorGreen(cell)
	case "关注":
		NewTableSet().SetBackgroundColorYellow(cell)
	case "警告":
		NewTableSet().SetBackgroundColorOrange(cell)
	case "危急":
		NewTableSet().SetBackgroundColorRed(cell)
	}
}

func AddPara(cell *document.Cell) *document.Paragraph {
	para := cell.AddParagraph()
	para.Properties().Spacing().SetLineSpacing(17, wml.ST_LineSpacingRuleAuto)
	para.Properties().SetAlignment(wml.ST_JcCenter)
	return &para
}

func AddParaLeft(cell *document.Cell) *document.Paragraph {
	para := cell.AddParagraph()
	para.Properties().Spacing().SetLineSpacing(17, wml.ST_LineSpacingRuleAuto)
	return &para
}

func AddRun(para *document.Paragraph, text string, size int, isBold bool) {
	run := para.AddRun()
	run.Properties().SetCharacterSpacing(0.05)
	run.Properties().SetBold(isBold)
	run.Properties().SetFontFamily("宋体")
	run.Properties().SetSize(measurement.Distance(measurement.Point * size))
	run.AddText(text)
}
