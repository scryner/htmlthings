package htmlthings

import (
	"code.google.com/p/go.net/html"
	"io"
)

type HtmlThings struct {
	docRoot *html.Node
}

func NewHtmlThings(r io.Reader) (*HtmlThings, error) {
	docRoot, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	return &HtmlThings{
		docRoot: docRoot,
	}, nil
}

func (things *HtmlThings) FindAttrValues(cond *Condition, attrName string) (ret []string, err error) {
	return
}

func (things *HtmlThings) FindTexts(cond *Condition) (ret []string, err error) {
	return
}
