package text

import (
	"github.com/carmel/gooxml/document"
	"github.com/carmel/gooxml/measurement"
	"github.com/carmel/gooxml/schema/soo/wml"
)

func AddPara(doc *document.Document) *document.Paragraph {
	para := doc.AddParagraph()
	para.Properties().Spacing().SetLineSpacing(24.5, wml.ST_LineSpacingRuleAtLeast)
	para.Properties().SetAlignment(wml.ST_JcBoth)
	return &para
}

func AddHeading1(doc *document.Document, text string) {
	para := doc.AddParagraph()
	para.Properties().Spacing().SetLineSpacing(40, wml.ST_LineSpacingRuleAtLeast)
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run := para.AddRun()
	run.Properties().SetCharacterSpacing(0.05)
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("宋体")
	run.Properties().SetSize(22)
	run.AddText(text)
}

func AddHeading2(doc *document.Document, text string) {
	para := doc.AddParagraph()
	para.Properties().Spacing().SetLineSpacing(40, wml.ST_LineSpacingRuleAtLeast)
	run := para.AddRun()
	run.Properties().SetCharacterSpacing(0.05)
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("宋体")
	run.Properties().SetSize(18)
	run.AddText(text)
}

func AddHeading3(doc *document.Document, text string) {
	para := doc.AddParagraph()
	para.Properties().SetStartIndent(32)
	para.Properties().Spacing().SetLineSpacing(40, wml.ST_LineSpacingRuleAtLeast)
	run := para.AddRun()
	run.Properties().SetCharacterSpacing(0.05)
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("宋体")
	run.Properties().SetSize(16)
	run.AddText(text)
}

func AddPageBreak(doc *document.Document) {
	doc.AddParagraph().AddRun().AddPageBreak()
}

func AddBreak(para *document.Paragraph) {
	para.AddRun().AddBreak()
}

func AddRun(para *document.Paragraph, text string, size int, isBold bool) {
	run := para.AddRun()
	run.Properties().SetCharacterSpacing(0.05)
	run.Properties().SetBold(isBold)
	run.Properties().SetFontFamily("宋体")
	run.Properties().SetSize(measurement.Distance(measurement.Point * size))
	run.AddText(text)
}
